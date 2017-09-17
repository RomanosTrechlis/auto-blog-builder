package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/RomanosTrechlis/blog-generator/cli"
	"github.com/RomanosTrechlis/blog-generator/config"
	"github.com/RomanosTrechlis/blog-generator/util/fs"
	"golang.org/x/net/context"
	"golang.org/x/tools/godoc/redirect"
	"google.golang.org/grpc"

	pb "github.com/RomanosTrechlis/auto-blog-builder/service"
)

const (
	fetchService    = "172.17.0.2:50051"
	generateService = "172.17.0.3:50051"
	defaultName     = "world"
	configFile      = "../artifacts/config.json"
)

var builder blogAutoBuilder

type blogAutoBuilder struct {
	Lock     sync.RWMutex
	SiteInfo *config.SiteInformation
}

var interval time.Duration

func init() {
	flag.DurationVar(&interval, "minutes", 600, "time between builds")
}

func main() {
	builder = blogAutoBuilder{
		SiteInfo: cli.ReadConfig(configFile),
	}
	fs.CreateFolderIfNotExist(builder.SiteInfo.DestFolder)

	go backgroundBuilder()

	http.HandleFunc("/", website)
	http.HandleFunc("/fetch/", fetch)
	http.HandleFunc("/generate/", generate)

	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)

}
func fetch(writer http.ResponseWriter, request *http.Request) {
	builder.Lock.Lock()
	defer builder.Lock.Unlock()
	downloader()
	writer.Write([]byte("Resources downloaded."))
}
func generate(writer http.ResponseWriter, request *http.Request) {
	builder.Lock.Lock()
	defer builder.Lock.Unlock()
	builderFunc()
	writer.Write([]byte("Website built"))
	redirect.Handler("/")
}

func website(writer http.ResponseWriter, request *http.Request) {
	l := len(request.URL.Path)
	if l < 1 {
		l = 1
	}
	file := fmt.Sprintf("%s/%s", builder.SiteInfo.DestFolder, request.URL.Path[1:l])
	fmt.Println(file)
	http.ServeFile(writer, request, file)
}

func backgroundBuilder() {
	for {
		builder.Lock.Lock()
		downloader()
		builderFunc()
		builder.Lock.Unlock()

		time.Sleep(interval * time.Minute)
	}
}

func downloader() {
	type source struct {
		dsType, dsRepo, dsDest string
	}
	datasources := []source{
		{
			dsType: builder.SiteInfo.DataSource.Type,
			dsRepo: builder.SiteInfo.DataSource.Repository,
			dsDest: builder.SiteInfo.TempFolder,
		},
		{
			dsType: builder.SiteInfo.Theme.Type,
			dsRepo: builder.SiteInfo.Theme.Repository,
			dsDest: builder.SiteInfo.ThemeFolder,
		},
	}
	services := []string{
		fetchService,
	}

	for _, s := range datasources {
		// Set up a connection to the server.
		conn, err := grpc.Dial(services[0], grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewFetcherClient(conn)
		req := &pb.FetchRequest{
			DsType:       s.dsType,
			DsRepository: s.dsRepo,
			DsDestFolder: s.dsDest,
		}
		r, err := c.Fetch(context.Background(), req)
		if err != nil {
			log.Fatalf("failled: %v", err)
		}
		log.Printf("Response: %s", r.Res)
	}
}

func builderFunc() {
	conn, err := grpc.Dial(generateService, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGeneratorClient(conn)
	r, err := c.Generate(context.Background(), pb.MapSiteInformationToRequest(builder.SiteInfo))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Res)
}
