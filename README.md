# OWL

An ease to use directory and/or file watcher. It keeps track of changes in the file or directory specified and sends notifications of changes to the resource. As of know it uses a short `polling` approach where it checks the resource **stats** and sends the notifications based on that. Moving forward I plan on adding a `system` call mode where it work using system events to check for changes.

**Side Note**
`I may or may not fully commit to system mode if I find it trully beneficial that may be determined by performance and scalability`

### Why Owl ?

I need a tool/lib like this for a project I'm working on. I did try to use a different library. However, It wasn't working as I expected it to. It could be skill issues or a case of `RTFM`, regardless it got me thinking how hard could it be to implement my own with just the features I need instead of an entire library of which I may use only 20% and so `OWL` was born.

I'm also interested in learning from the process of putting it to together. That is mainly the reason why I want to try and implement a system call version where I rely on events from the kernel to identify changes in resources. I'm not sure at this point if that is possible but I beleieve it is.

### Usage

At this point usage is very simple. Do keep in mind this is not an stable release therefore the API may change in the future.

```go
// initialize the watcher
w, err := owl.NewWatcher("../sandbox", owl.Options{})
if err != nil {
  t.Error(err)
}
for {
  // handle events, errors and done channels
  select {
  case e, open := <-w.Events:
    if !open {
      return
    }
    if e.Operation == owl.FileModified {
      fmt.Printf("File was modified. %s\n", e.Location)
    } else {
      fmt.Printf("%v Something else happend to the file\n", e)
    }
  case err, open := <-w.Errors:
    if !open {
      return
    }
    t.Error(err)
  case _ = <-w.Done:
    fmt.Println("Process is done")
  }
}
```

