# Keakie Platform Coding Challenge
Welcome to the Keakie Platform coding challenge. This is designed to demonstrate your knowledge and understanding of the Go programming langauge.

## TL;DR
- Create two small microservices using [`Go kit`](https://github.com/go-kit/kit), which communicate with each other using RPC (Remote Procedure Call)
- Please don't spend more than 3-4 hours on this challenge. Your time is valuable
- **Good luck!**

## Background Reading
Keakie is a curated audio streaming platform, primarily consisting of a large selection of "shows" that a listener can explore:

- `show`: A collection of episodes, presented as a branded experience, with a host (DJ), title, description, genres, imagery and other metadata. A show can have multiple genres

- `genre`: A genre of music, which in turn can be related to other genres

## Requirements

**GOAL**
- **Create two small microservices - one for shows and one for genres**
- **Be able to fetch (via an API call) a show, which includes an array of genre objects in the response**

### Genres microservice
- We can pass a genre "slug", from the list in `test-genre-slugs.txt`, to the service. Eg `/genres/{slug}`
- For a valid slug, the service retrieves the appropriate genre from the DB (see `db-mocks/genres-db-mock.json`)
- The service endpoint responds with the genre object (as JSON - see `example-response-body/genre-by-slug.json`)

### Shows microservice
- We can pass a show "slug", from the list in `test-show-slugs.txt`, to the service. Eg `/shows/{slug}`
- For a valid slug, the service retrieves the appropriate show from the DB (see `db-mocks/shows-db-mock.json`)
- The service replaces the genres array [of slugs] with a respective array of genre objects, retrieved from the genres microservice via RPC
- The service endpoint responds with the show object (as JSON - see `example-response-body/show-by-slug.json`)

### Essentials
- Use [`Go kit`](https://github.com/go-kit/kit) and its guidelines as best you can
- The shows microservice must communicate with the genres microservice via RPC
- We don't require you to use any Database technology - please use the mock DB responses from `db-mocks`. Try to organise the code in a way that would make it easy to switch to an actual DB in the future
- Place all your code in the `src` directory
- Provide instructions on how to run the services locally, and/or any other related instructions/notes in `src/README.md`
- Clone the repository initially, and then provide us with a link to a repo containing your completed challenge

### Bonus tasks
- Basic error handling
- Basic testing suite

### Tips
- Reference the [Go kit examples](https://github.com/go-kit/examples)
- Focus on the most important tasks first. At a minimum, create a shows microservice which can be called to exchange a show slug for a show object
- We care about quality over quantity. If you start to exceed 3-4 hours, please submit what you have completed so far and add notes to `src/README.md`
