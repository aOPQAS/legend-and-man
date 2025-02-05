package main

func main() {
	speaker := &Speaker{}

	oncommand := &OnCommand{
		device: speaker,
	}

	offcommand := &OffCommand{
		device: speaker,
	}

	onbutton := &Button{
		command: oncommand,
	}
	onbutton.press()

	offbutton := &Button{
		command: offcommand,
	}

	offbutton.press()
}
