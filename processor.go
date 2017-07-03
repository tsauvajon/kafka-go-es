package main

func (e CreateEvent) Process() error {
	return nil
}

func (e InvalidEvent) Process() error {
	return nil
}

func (e AcceptEvent) Process() error {
	return nil
}
