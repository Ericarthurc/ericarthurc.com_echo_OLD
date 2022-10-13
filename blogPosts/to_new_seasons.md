---
title: To New Seasons
date: April 9, 2022
tags: new season, next gen, golang, echo, blog
series: blog
---

# To New Seasons

The site is finally live! This is my personal blog/thoughts collector. I wont ever write anything really personal here, as this is still the internet, but this will be more to talk about my thoughts on all things Tech around the world. The website will be in a beta state for a long time to come; I have a lot of ideas for the site, and it will take time to bring them together. It's like an evolving code project to show off my knowledge and learn.

The core of the site was already in place before I started designing too. I was really getting into backend server markdown parsers, and I started developing different backend rendered web sites around this concept. I have written many different builds in JS, TS, Nim, Goland and Rust. And I have bounced back and forth between tons of builds and frameworks and ideas... I have an issue with settling 😛.

Rough idea of what I am looking for in a project language:

- Fast HTTP response times
- Good community support
- Well developed packages that support the project
- Semi-quick development
- Fun development 🎉

Golang just seems to fix this mold the best; I have written a lot in it as well so I feel comfortable and it's quick to pick up new ideas in and debug. Rust is a beautiful language with some amazing concepts, but in the HTTP space the performance is about the same vs Golang; which then makes it hard to justify the massive development time increase. Rust just makes you feel smart, it does things 'right' and forces you to aswell... it just takes a long time to develop in. And when learning new concepts in Rust you are left trying to wrap your mind around these huge boilerplate concepts and large complex types. I just want to write code, see it come to life and have fun! If you're looking for a good Rust HTTP server, start with Axum! It is probably the best Rust server you can use right now! Great project by the Tokio team.

Then Rust, which is a beautiful language with some amazing concepts! Rust just makes you feel smart haha; it does things 'right' and forces you to aswell... it just can take a long time to develop in it though... And when learning new concepts in Rust you are left trying to wrap your mind around these huge boilerplate concepts and large complex types. This is a very good language with a lot of potential; I just don't want to spent months writing code. I love Rust, and use it for other things but for this project it's going to sit in the back corner... for now! Haha!

Lets look at the JS/TS world real quick, which I hate the state of it right now. Give it a couple more years and it will be in a perfect state. Nodejs should probably just die at this point, and I am only slightly joking. Node is just missing a lot when it comes to a user friendly modern JS/TS experience; go try and use top level async/await with TypeScript on a Node project... good luck! And then Deno offers great modern tooling, but performance is still behind and packages are minimal. Deno has a great potential but is just too young; they are having a hard time convincing people to move over from Node and loose 95% of the mature packages they use everyday. And then you have Bun, which is super fast but again very very early and just not ready yet. Give it a couple years and I think Deno and Bun take off. Node will probably be around forever just to maintain legacy code; but a lot of developers have jumped ship to Golang/Rust already.

### Framework

I will be using the <a href="https://github.com/labstack/echo" target="_blank" rel="noreferrer noopener">Echo</a> package for this project, as I like this framework compared to some other designs. This package is nice and simple; they pride themselves on being minimalist which is great. Echo is also built on top of `http/net` which is another big plus; this keeps the package smaller and more secure. `http/net` is maintained by the developers of Go so it will also be true to the nature of the language and managed by really smart people.

So here is a little snippet of a Echo HTTP Server:

```go
package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
)

func main() {
  	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
```

So what is next?! Well hopefully I keep developing this site and learning new things to implement. And I hope the content is good and enjoyable!

### Eric Christensen
