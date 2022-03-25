osu
===
`go mod init github.com/makinap/osu`

`go mod vendor`

`docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres`

`docker build --tag docker-gs-ping .`

`docker run --publish 1323:1323 docker-gs-ping`

`goose status`

`goose up`

`gqlgen`



Graph examples
```
{
  tasks{
    id, title, note
  }
}
```

```
mutation {
  createTask(input: {
      title: "ee",
      note: "aa"
    }) {
      title
      note
  }
}
```

```
mutation {
  auth{
    register(input: {
      name: "my name",
      email: "tineo@example.com",
      password: "1234"
    })
  }
}
```

```
mutation {
  auth{
   login(email: "tineo@example.com",password: "1234")
  }
}
```