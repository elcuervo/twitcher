# Twitcher

The ticker updater

```go
twitch := twitcher.NewTwitcher(time.Seconds * 5)
twitcher.Do(func() {
        print("I'm going to be executed every 5 seconds\n")
})
```

Update the execution time:
```go
twitcher.Update(time.Hour * 2)
```

You can also access the channel without fancy stuff:

```go
for {
        <-twitcher.C
}
```
