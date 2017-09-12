// go pprof1 OMIT
func main() {
  ....
  f, err := os.Create("cpu-profile.prof")
  if err != nil {
    log.Fatal(err)
  }
  pprof.StartCPUProfile(f)
  ... // this is program you want to profile
  pprof.StopCPUProfile()
}

// go pprof2 OMIT
