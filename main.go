package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/MatusOllah/slogcolor"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func getLogLevel(s string) slog.Leveler {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func showVersion() {
	a := app.New()

	fmt.Printf("%s version v%s\n", a.Metadata().Name, a.Metadata().Version)
	fmt.Printf("Go version %s\n", runtime.Version())
}

func main() {
	path := flag.String("open-file", "", "Open a file")
	logLevel := flag.String("log-level", "info", "Log level (\"debug\", \"info\", \"warn\", \"error\")")
	sr := flag.Int("sample-rate", 8000, "Project Sample Rate")
	_mode := flag.String("mode", "Bytebeat", "The mode (\"Bytebeat\", \"Floatbeat\")")
	version := flag.Bool("version", false, "Show version and exit")
	flag.Parse()

	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, &slogcolor.Options{
		Level:       getLogLevel(*logLevel),
		TimeFormat:  time.DateTime,
		SrcFileMode: slogcolor.ShortFile,
	})))
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	if *version {
		showVersion()
		os.Exit(0)
	}

	slog.Info("Initializing")
	beforeInit := time.Now()

	a := app.New()

	a.Lifecycle().SetOnStarted(func() {
		slog.Info("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		slog.Info("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		slog.Info("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		slog.Info("Lifecycle: Exited Foreground")
	})

	slog.Info(fmt.Sprintf("gecfg-editor version %s", a.Metadata().Version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))

	slog.Info("initializing speaker")
	sampleRate = beep.SampleRate(a.Preferences().IntWithFallback("OutputSampleRate", 48000))
	slog.Info("using output sample rate", "sampleRate", sampleRate)
	speaker.Init(sampleRate, 1024)

	bbSampleRate = beep.SampleRate(*sr)

	if *_mode != "Bytebeat" && *_mode != "Floatbeat" {
		panic("invalid mode: " + *_mode)
	}
	mode = *_mode

	w := a.NewWindow(openFileName + " - " + a.Metadata().Name)
	w.SetMaster()
	w.Resize(fyne.NewSize(1280, 720))
	w.SetMainMenu(makeMainMenu(a, w))

	w.SetContent(makeUI(a, w))

	slog.Info(fmt.Sprintf("Initialization took %s", time.Since(beforeInit)))

	// open file
	if *path != "" {
		slog.Info("opening file", "path", *path)

		content, err := os.ReadFile(*path)
		if err != nil {
			slog.Error(err.Error())
			panic(err)
		}
		expr := string(content)

		openFileName = filepath.Base(*path)
		openFilePath = *path
		exprEntry.SetText(expr)
		if err := bbgen.SetExpr(expr); err != nil {
			slog.Error(err.Error())
			panic(err)
		}
		updateWindowTitle(a, w)
	}

	w.ShowAndRun()

	slog.Info("exiting")
	os.Exit(0)
}
