package main

import "github.com/nsf/termbox-go"
import "math/rand"
import "time"

func draw() { // 随机设置字符单元的属性
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault,
				termbox.Attribute(rand.Int()%8)+1)
		}
	}
	termbox.Flush() // 刷新后台缓存到界面中
}

func main() {
	err := termbox.Init() // 初始化termbox包
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent() // 开始监听键盘事件
		}
	}()

	draw()
	for {
		select {
		case ev := <-event_queue: // 获取键盘事件
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				return //如果是 Esc 键，则退出程序
			}
		default:
			draw()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
