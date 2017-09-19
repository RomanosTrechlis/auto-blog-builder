## auto-blog-builder
Playing with docker containers and service using as example the blog-generator

### 1. Service

The service folder contains the proto buf service definition.  There are two services:

1. the fetch service, and
1. the generate service.

The first retrieves the necessary artifacts and the second builds the blog.

### 2. Fetch

The server that retrieves artifacts like the blog's theme and the actual blog posts.

### 3. Generate

The server that generates the static pages of the blog.

### 4. Web

A client that requests the retrieval and generation of the blog, as well as a basic file server to display the generated static pages.

## TODO

Use of a service discovery middleware.