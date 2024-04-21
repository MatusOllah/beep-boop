GO = go
FYNE = $(GO) run fyne.io/fyne/v2/cmd/fyne@latest

TARGET=$(shell $(GO) env GOOS)
BINARY=./bin/$(TARGET)

all: build
.PHONY: all

build: clean
build:
	mkdir -p $(BINARY)
	$(GO) get
	$(FYNE) build --target $(TARGET) -o $(BINARY)
.PHONY: build

package: clean
	$(GO) get
	$(WINRES) make
	$(FYNE) package --target $(TARGET)
.PHONY: package

clean:
	rm -rf $(BINARY)
.PHONY: build
