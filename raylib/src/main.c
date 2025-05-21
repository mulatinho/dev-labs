#include "raylib.h"
#include <stdlib.h>
#include <time.h>
#include <unistd.h>

#define TILE_WIDTH 64
#define TILE_HEIGHT 32
#define MAP_WIDTH 32
#define MAP_HEIGHT 32

typedef struct {
    int x, y;
} Player;

// Converts grid position to screen position in isometric
Vector2 GridToScreen(int x, int y) {
    return (Vector2){
        (x - y) * TILE_WIDTH / 2 + 800 / 2,  // X center offset
        (x + y) * TILE_HEIGHT / 2 + 50       // Y offset
    };
}

int main(void) {
    InitWindow(1024, 1024, "Isometric Map with Tileset");
    SetTargetFPS(60);

    // Load textures
    Texture2D tileset = LoadTexture("assets/tile-beach-sand-sm.png");
    Texture2D character = LoadTexture("assets/castaway-man.png");

    // Define tileset info
    int tile_count = 2;
    Rectangle tileRects[2];
    for (int i = 0; i < tile_count; i++) {
        tileRects[i] = (Rectangle){ 0, 0, TILE_WIDTH, TILE_HEIGHT };
    }

    // Generate map
    int map[MAP_HEIGHT][MAP_WIDTH];
    srand(time(NULL));
    for (int y = 0; y < MAP_HEIGHT; y++) {
        for (int x = 0; x < MAP_WIDTH; x++) {
            map[y][x] = rand() % tile_count; // Random tile 0-3
        }
    }

    Player player = {16, 16};

    while (!WindowShouldClose()) {
        // Input
        if (IsKeyPressed(KEY_RIGHT)) player.x++;
        if (IsKeyPressed(KEY_LEFT))  player.x--;
        if (IsKeyPressed(KEY_DOWN))  player.y++;
        if (IsKeyPressed(KEY_UP))    player.y--;
        if (IsKeyPressed(KEY_Q))     goto cleanup;

        BeginDrawing();
        ClearBackground(WHITE);

        // Draw tilemap
        for (int y = 0; y < MAP_HEIGHT; y++) {
            for (int x = 0; x < MAP_WIDTH; x++) {
                int tileId = map[y][x];
                Vector2 pos = GridToScreen(x, y);
                DrawTextureRec(tileset, tileRects[tileId], pos, WHITE);
            }
        }

        // Draw player sprite
        Vector2 playerPos = GridToScreen(player.x, player.y);
   	DrawTextureV(character, (Vector2){ playerPos.x - 16, playerPos.y - 48 }, WHITE);
	//DrawPoly((Vector2){playerPos.x, playerPos.y}, 4, TILE_WIDTH / 2, 45.0f, RED);

        DrawText("Use arrow keys to move", 10, 10, 20, RAYWHITE);
        EndDrawing();

	usleep(4500);
    }

cleanup:
    UnloadTexture(tileset);
    UnloadTexture(character);
    CloseWindow();
    return 0;
}
