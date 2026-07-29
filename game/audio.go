package game

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

//go:embed assets/music/village.ogg
var villageMusic []byte

func (g *Game) PlaySong() error {
	stream, err := vorbis.DecodeF32(bytes.NewReader(villageMusic))
	if err != nil {
		return fmt.Errorf("decode music: %w", err)
	}

	audioContext := audio.NewContext(stream.SampleRate())

	loop := audio.NewInfiniteLoopF32(
		stream,
		stream.Length(),
	)

	musicPlayer, err := audioContext.NewPlayerF32(loop)
	if err != nil {
		return fmt.Errorf("create music player: %w", err)
	}

	musicPlayer.SetVolume(0.35)
	musicPlayer.Play()

	g.audioContext = audioContext
	g.musicPlayer = musicPlayer

	return nil
}
