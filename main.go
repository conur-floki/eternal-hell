package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	Position  rl.Vector2
	Speed     float32
	Active    bool
	Direction rl.Vector2
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(800)

	rl.InitWindow(screenWidth, screenHeight, "Eternal hell")

	playerPosition := rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	playerSpeed := float32(5)

	bulletsMax := 30
	bullets := make([]Bullet, bulletsMax)

	for i := 0; i < bulletsMax; i++ {
		bullets[i] = Bullet{Position: rl.NewVector2(0, 0), Speed: 10, Active: false}
	}

	rl.SetTargetFPS(60)
	backgroundTexture := rl.LoadTexture("./assets/pastito.png")
	playerSheet := rl.LoadTexture("./assets/necro.png")
	frameWidth := float32(128)
	frameHeight := float32(128)
	playerPosition = rl.NewVector2(100, float32(screenHeight/2))
	frameCounter := 0
	animationFrames := []int{8, 13, 17, 5, 10}

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyW) && playerPosition.Y > 35 { // Check upper limit
			playerPosition.Y -= playerSpeed
		}
		if rl.IsKeyDown(rl.KeyS) && playerPosition.Y < float32(screenHeight-35) { // Check lower limit
			playerPosition.Y += playerSpeed
		}
		if rl.IsKeyDown(rl.KeyA) && playerPosition.X > 35 { // Check left limit
			playerPosition.X -= playerSpeed
		}
		if rl.IsKeyDown(rl.KeyD) && playerPosition.X < float32(screenWidth-35) { // Check right limit
			playerPosition.X += playerSpeed
		}

		if rl.IsKeyPressed(rl.KeyUp) {
			fireBullet(&bullets, playerPosition, rl.NewVector2(0, -1))
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			fireBullet(&bullets, playerPosition, rl.NewVector2(0, 1))
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			fireBullet(&bullets, playerPosition, rl.NewVector2(-1, 0))
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			fireBullet(&bullets, playerPosition, rl.NewVector2(1, 0))
		}

		for i := 0; i < bulletsMax; i++ {
			if bullets[i].Active {
				bullets[i].Position.X += bullets[i].Speed * bullets[i].Direction.X
				bullets[i].Position.Y += bullets[i].Speed * bullets[i].Direction.Y

				if bullets[i].Position.Y < 0 || bullets[i].Position.Y > float32(screenHeight) ||
					bullets[i].Position.X < 0 || bullets[i].Position.X > float32(screenWidth) {
					bullets[i].Active = false
				}
			}
		}

		rl.BeginDrawing()

		rl.DrawTexturePro(
			backgroundTexture,
			rl.NewRectangle(0, 0, float32(backgroundTexture.Width), float32(backgroundTexture.Height)),
			rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight)),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)

		for i := 0; i < bulletsMax; i++ {
			if bullets[i].Active {
				rl.DrawCircleV(bullets[i].Position, 3, rl.Maroon)
			}
		}

		currentFrame := 0
		for i, frames := range animationFrames {
			if frameCounter >= currentFrame && frameCounter < currentFrame+frames {
				rl.DrawTextureRec(playerSheet, rl.NewRectangle(0, frameHeight*float32(i), frameWidth, frameHeight), playerPosition, rl.White)
				break
			}
			currentFrame += frames
		}

		frameCounter++
		if frameCounter >= currentFrame {
			frameCounter = 0
		}

		rl.EndDrawing()
	}

	rl.UnloadTexture(backgroundTexture)
	rl.CloseWindow()
}

func fireBullet(bullets *[]Bullet, position rl.Vector2, direction rl.Vector2) {
	for i := 0; i < len(*bullets); i++ {
		if !(*bullets)[i].Active {
			bulletSpawnPosition := rl.NewVector2(position.X+100, position.Y+75)

			(*bullets)[i].Active = true
			(*bullets)[i].Position = bulletSpawnPosition // Usar la nueva posición del disparo
			(*bullets)[i].Direction = direction
			break
		}
	}
}
