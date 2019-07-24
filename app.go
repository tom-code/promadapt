package main

import (
  "fmt"
  "net/http"
  "time"


  
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
  "github.com/prometheus/client_golang/prometheus/promhttp"

)

func main() {
  fmt.Println("start")

  var cnt1 float64
  var cnt2 float64
  var cnt3 float64

  cnt2 = 1000000
  cnt3 = 2000000

  promauto.NewCounterFunc(prometheus.CounterOpts{
    Name: "testapp_cnt_1",
  },
  func() float64 {
    return float64(cnt1)
  })

  promauto.NewCounterFunc(prometheus.CounterOpts{
    Name: "testapp_cnt_multi",
    ConstLabels: prometheus.Labels{ "label1": "aaaa" },
  },
  func() float64 {
    return float64(cnt2)
  })

  promauto.NewCounterFunc(prometheus.CounterOpts{
    Name: "testapp_cnt_multi",
    ConstLabels: prometheus.Labels{ "label1": "bbbb" },
  },
  func() float64 {
    return float64(cnt3)
  })



  http.Handle("/metrics", promhttp.Handler())
  go http.ListenAndServe(":9999", nil)

  for {
    cnt1 = cnt1 + 1
    time.Sleep(1*time.Second)
  }

}
