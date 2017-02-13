package main

import (
    "github.com/warthog618/gpio"
    "github.com/alexflint/go-arg"
)

var args struct {
    State     string `arg:"positional,help:off | on"`
    Devices []string `arg:"positional,help:led | pi1 | pi2 | pi3 | pi4"`
}

var pin_nums = map[string] uint8 {
    "led": gpio.J8_29,
    "pi1": gpio.J8_31,
    "pi2": gpio.J8_33,
    "pi3": gpio.J8_35,
    "pi4": gpio.J8_37,
}

func main() {
    arg.MustParse(&args)

    err := gpio.Open()
    if err != nil {
        panic(err)
    }
    defer gpio.Close()

    level := gpio.Low
    if args.State == "on" {
        level = gpio.High
    }

    for _, dev := range args.Devices {
        pin_num, ok := pin_nums[dev]
        if ok {
            pin := gpio.NewPin(pin_num)
            pin.Output()
            pin.Write(level)
        }
    }
}
