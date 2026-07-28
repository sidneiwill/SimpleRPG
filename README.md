# RPG Movement Prototype

A small top-down RPG movement prototype built with Go and Ebiten. The project focuses on the fundamentals behind classic 2D RPG exploration.

This is a learning project, but it is being organized with real game-development structure in mind: clear packages, small systems, and code that can grow using correct structures and scalability. 

The prototype can grow into a more complete RPG foundation over time, with features like map collisions, NPCs, dialogue boxes, interactable objects, scene transitions, quests, enemy encounters, and simple save/load support.

![Sprout Lands asset pack preview](https://img.itch.zone/aW1nLzc2NDY4MDEuZ2lm/original/9zxkqt.gif)

Created using assets from [Sprout Lands - Asset Pack](https://cupnooble.itch.io/sprout-lands-asset-pack) by Cup Nooble.

## What It Does

- Loads a character sprite sheet.
- Animates walking in multiple directions.
- Draws a larger map image as the world background.
- Moves the player with `W`, `A`, `S`, and `D`.
- Uses a camera to follow the player across the map.
- Opens and closes a simple inventory menu with `I`.
- Lets the player move the inventory cursor with `W` and `S`.
- Pauses player movement while the inventory menu is open.
- Keeps rendering at a fixed internal resolution while scaling the game window.

## Tech Stack

- Go
- Ebiten
- Embedded PNG assets with `go:embed`

## Assets & Credits

This project uses free assets from [Sprout Lands - Asset Pack by Cup Nooble](https://cupnooble.itch.io/sprout-lands-asset-pack) on itch.io.

The free version is intended for non-commercial projects, according to the asset page. Credit goes to Cup Nooble.

## Running The Project

Install dependencies and run the game:

```sh
go run .
```

Run checks:

```sh
go test ./...
```

## Controls

```text
W - move up
A - move left
S - move down
D - move right
I - open/close inventory

When inventory is open:
W - move cursor up
S - move cursor down
```

## Current State

The project is currently split into small packages and files:

```text
main.go              app startup only
game/game.go         Game struct, Update, Draw, Layout
game/camera.go       camera movement
game/assets.go       asset loading
game/inventory.go    inventory input and drawing
game/player_movement.go player movement input
player/player.go     player animation state
sprite/manager.go    reusable sprite animation system
internal/mathutil    shared math helpers
```

The current inventory is an early menu overlay. It displays a small item list, tracks a cursor, wraps the cursor at the top and bottom, and stops player movement while open.

## Roadmap

### Gameplay Features

- [ ] Add a collision system.
  - Prevent the player from walking through walls, trees, buildings, water, and other blocked map areas.
  - Start simple with rectangle-based collision zones before moving to tile-based collision.

- [ ] Add text/dialogue boxes.
  - Display RPG-style message windows for NPC conversations, signs, item descriptions, and story events.
  - Support progressive text rendering so dialogue appears naturally instead of all at once.

- [x] Add a first inventory overlay.
  - Open and close the menu with `I`.
  - Move through the item list with `W` and `S`.
  - Pause movement while the menu is open.

- [ ] Expand the inventory into a fuller RPG menu.
  - Add categories such as items, status, equipment, save, and settings.
  - Add item selection/use behavior instead of cursor movement only.
  - Replace debug text with a proper font and UI style.

### Refactor Organization

- [ ] Extract drawing helpers if `Draw` keeps growing.
  - Good first shape: `g.drawMap(screen)` and `g.drawPlayer(screen)` in `game/render.go`.
  - This keeps `Draw` readable without creating a renderer abstraction too early.

- [ ] Tighten game state flow.
  - Move menu-state handling before world updates if the menu should fully pause animation and camera updates.
  - Consider an explicit game mode enum when adding dialogue, battle, or scene transitions.

- [ ] Introduce a small position type when more objects need coordinates.
  - Example: `type Vec2 struct { X, Y float64 }`.
  - Then replace pairs like `playerX`, `playerY` with `playerPosition Vec2`.
  - Skip this until there is more than one moving entity.

### Cleanup Notes

- [ ] Keep package names simple and consistent with folder names.
- [ ] Run `gofmt` after moving code.
- [ ] Run `go test ./...` after each refactor step.

### Testing

- [ ] Add application tests.
  - Cover core behavior such as player movement, camera following, collision rules, and menu state.
  - Start with small unit tests for pure logic before testing larger game flows.

## Why This Project Exists

This project is a practical study of the small systems that make a 2D RPG feel good: movement, animation timing, camera behavior, world rendering, and code organization. The focus is not just getting pixels on the screen, but learning how to structure a game loop so new features can be added cleanly.
