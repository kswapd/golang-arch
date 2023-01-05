package base

import (
	"flag"

	"k8s.io/klog/v2"
)

func KlogTest() {

	klog.Info("nice to meet dddd")
	klog.InitFlags(nil)
	if err := flag.Set("logtostderr", "false"); err != nil {
		panic(err)
	}
	if err := flag.Set("alsologtostderr", "false"); err != nil {
		panic(err)
	}
	if err := flag.Set("stderrthreshold", "fatal"); err != nil {
		panic(err)
	}
	if err := flag.Set("v", "0"); err != nil {
		panic(err)
	}
	flag.Parse()

	//buf := new(bytes.Buffer)
	//klog.SetOutput(buf)
	klog.Info("nice to meet you1")
	klog.Error("nice to meet you2")
	//klog.Fatal("nice to meet you3")
	//klog.Flush()

	//fmt.Printf("LOGGED: %s", buf.String())
}

func KlogTest2() {

	klog.Info("nice to meet eee")
	var vflag flag.FlagSet
	klog.InitFlags(&vflag)
	if err := vflag.Set("logtostderr", "false"); err != nil {
		panic(err)
	}
	if err := vflag.Set("alsologtostderr", "false"); err != nil {
		panic(err)
	}
	if err := vflag.Set("stderrthreshold", "fatal"); err != nil {
		panic(err)
	}
	if err := vflag.Set("v", "0"); err != nil {
		panic(err)
	}
	//vflag.Parse()

	//buf := new(bytes.Buffer)
	//klog.SetOutput(buf)
	klog.Info("nice to meet you1")
	klog.Error("nice to meet you2")
	//klog.Fatal("nice to meet you3")
	//klog.Flush()

	//fmt.Printf("LOGGED: %s", buf.String())
}
