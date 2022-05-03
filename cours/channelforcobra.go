package cours

import (
	"fmt"
	"github.com/mrlaowilly/devgo/internal/fssync"
	"github.com/spf13/cobra"
	"time"
)

func Mastercobra() {
	// creer un iPC (Internal Process Channel)
	c := make(chan string)
	copie := make(chan string)

	// 4 actions longue a effectuer
	go workercobra(c, func() string { return tachecobra(4) })
	go workercobra(c, func() string { return tachecobra(1) })
	go workercobra(c, func() string { return tachecobra(3) })
	go workercopie(copie, func() string { return tachecopie(20) })

	// recupère les resultat
	i := 0
	for result := range c {
		fmt.Println(result)

		if i >= 2 {
			break
		}
		i++
	}
	u := 0
	for result := range copie {
		fmt.Println(result)

		if u >= 1 {
			break
		}
		u++
	}
}

func tachecobra(delta time.Duration) string {
	time.Sleep(delta * time.Second)
	return fmt.Sprintf("%v", delta*time.Second)
}

func workercobra(c chan string, callback func() string) {
	c <- callback()
}

func workercopie(copie chan string, tachecopie func() string) {
	copie <- tachecopie()
}

func tachecopie(delta time.Duration) string {
	var root = &cobra.Command{
		Use:   "app [source] [target]",
		Short: "Permet de copier des fichiers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			target := args[1]

			if fssync.FolderExist(source) {
				fssync.CreatFolderIfNotExist(target)
				for _, file := range fssync.Scan(source, source) {
					fssync.Copy(source+file, target+file)
				}
			} else {
				fmt.Println("Source folder not exist")
			}

			fmt.Println(source, target)
		},
		DisableFlagsInUseLine: true,
	}

	root.Execute()
	time.Sleep(delta * time.Second)
	return fmt.Sprintf("%v", delta*time.Second)

}
