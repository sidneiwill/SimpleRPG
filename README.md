# RPG Movement Prototype 🗺️

A small top-down RPG movement prototype built with Go and Ebiten. The project focuses on the fundamentals behind classic 2D RPG exploration.

This is a learning project, but it is being organized with real game-development structure in mind: clear packages, small systems, and code that can grow using correct structures and scalability. 

The prototype can grow into a more complete RPG foundation over time, with features like map collisions, NPCs, dialogue boxes, interactable objects, scene transitions, quests, inventory, enemy encounters, and simple save/load support.

![Sprout Lands asset pack preview](https://img.itch.zone/aW1nLzc2NDY4MDEuZ2lm/original/9zxkqt.gif)

Created using assets from [Sprout Lands - Asset Pack](https://cupnooble.itch.io/sprout-lands-asset-pack) by Cup Nooble.

## What It Does ✨

- Loads a character sprite sheet.
- Animates walking in multiple directions.
- Draws a larger map image as the world background.
- Moves the player with `W`, `A`, `S`, and `D`.
- Uses a camera to follow the player across the map.
- Keeps rendering at a fixed internal resolution while scaling the game window.

## Tech Stack 🧰

- Go
- Ebiten
- Embedded PNG assets with `go:embed`

## Assets & Credits 🎨

This project uses free assets from [Sprout Lands - Asset Pack by Cup Nooble](https://cupnooble.itch.io/sprout-lands-asset-pack) on itch.io.

The free version is intended for non-commercial projects, according to the asset page. Credit goes to Cup Nooble.

## Project Structure 📁

```text
.
├── assets/
│   ├── character.png
│   └── map.png
├── game/
│   └── camera.go
├── player/
│   └── player.go
├── sprite/
│   └── manager.go
├── main.go
├── go.mod
└── TODO.md
```

## Running The Project ▶️

Install dependencies and run the game:

```sh
go run .
```

Run checks:

```sh
go test ./...
```

## Controls 🎮

```text
W - move up
A - move left
S - move down
D - move right
```

## Current Direction 🧭

The next goal is to keep reducing the responsibilities inside `main.go`. Right now the project is moving toward this shape:

```text
main.go              app startup only
game/game.go         Game struct, Update, Draw, Layout
game/camera.go       camera movement
game/assets.go       asset loading
game/constants.go    screen and tuning constants
player/player.go     player animation state
sprite/manager.go    reusable sprite animation system
```

## Roadmap 🛠️

### Gameplay Features

- [ ] Add a collision system.
  - Prevent the player from walking through walls, trees, buildings, water, and other blocked map areas.
  - Start simple with rectangle-based collision zones before moving to tile-based collision.

- [ ] Add text/dialogue boxes.
  - Display RPG-style message windows for NPC conversations, signs, item descriptions, and story events.
  - Support progressive text rendering so dialogue appears naturally instead of all at once.

- [ ] Add an RPG Maker-style menu.
  - Create an in-game pause/menu screen with options like items, status, equipment, save, and settings.
  - Keep the first version simple: open/close the menu, move through options, and select an entry.

### Refactor Organization

- [ ] Rename the camera package or folder so they match.
  - Current state: `game/camera.go` says `package camera`, but it lives in the `game/` folder.
  - Better beginner option: move it to `camera/camera.go` and import `sidneiwill.dev/BaseRPGMovement/camera`.
  - Alternative: keep it in `game/`, but use `package game`.

- [ ] Extract the `Game` type and Ebiten methods from `main.go`.
  - Move `type Game`, `NewGame`, `Update`, `Draw`, and `Layout` into `game/game.go`.
  - Keep `main.go` responsible only for creating the game, setting the window, and calling `ebiten.RunGame`.

- [ ] Extract constants into a dedicated file.
  - Move `screenWidth`, `screenHeight`, `windowWidth`, `windowHeight`, and `cameraSmoothness` into `game/constants.go`.
  - If they must be used from `main.go`, export them as `ScreenWidth`, `WindowWidth`, etc.

- [ ] Extract asset loading from `main.go`.
  - Move `go:embed` variables and image loading code into an `assets.go` file.
  - Important: `go:embed` cannot use `../assets/map.png`, so either keep `assets.go` in the same folder as `assets/`, or move image files under the package that embeds them.

- [ ] Extract player movement/input handling.
  - Move the WASD logic out of `Update` once it grows more.
  - Good first shape: create a method like `g.updatePlayerMovement()` in `game/player_movement.go`.
  - Do not create a separate `input` package yet; a method is enough for now.

- [ ] Extract drawing helpers if `Draw` keeps growing.
  - Good first shape: `g.drawMap(screen)` and `g.drawPlayer(screen)` in `game/render.go`.
  - This keeps `Draw` readable without creating a renderer abstraction too early.

- [ ] Introduce a small position type when more objects need coordinates.
  - Example: `type Vec2 struct { X, Y float64 }`.
  - Then replace pairs like `playerX`, `playerY` with `playerPosition Vec2`.
  - Skip this until there is more than one moving entity.

### Cleanup Notes

- [ ] Remove duplicate `clamp` and `lerp` functions if they exist in more than one package.
- [ ] Keep package names simple and consistent with folder names.
- [ ] Run `gofmt` after moving code.
- [ ] Run `go test ./...` after each refactor step.

### Testing

- [ ] Add application tests.
  - Cover core behavior such as player movement, camera following, collision rules, and menu state.
  - Start with small unit tests for pure logic before testing larger game flows.

## Why This Project Exists

This project is a practical study of the small systems that make a 2D RPG feel good: movement, animation timing, camera behavior, world rendering, and code organization. The focus is not just getting pixels on the screen, but learning how to structure a game loop so new features can be added cleanly.
