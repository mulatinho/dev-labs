#include "raylib.h"
#include <stdio.h>

#define MAP_WIDTH 32
#define MAP_HEIGHT 32
#define TILE_SIZE 1.0f

int map[MAP_HEIGHT][MAP_WIDTH];

typedef struct {
	int map[MAP_HEIGHT][MAP_WIDTH];
	Model sandModel, treeModel;
	Camera3D camera;
} World;

typedef struct {
	Model pirateModel;
	Texture2D texPirate;
	int x;
	int y;
} Player;

Player player = {
	.x = MAP_WIDTH / 2, 
	.y = MAP_HEIGHT / 2
};

World world = {
	.camera = {
		.position = (Vector3){40.0f, 40.0f, 40.0f},
		   .target = (Vector3){16.0f, 0.0f, 16.0f},
		   .up = (Vector3){0.0f, 1.0f, 0.0f},
		   .fovy = 45.0f,
		   .projection = CAMERA_PERSPECTIVE
	}
};

void handle_input(void)
{
	if (IsKeyPressed(KEY_RIGHT)) {
		if (player.x < MAP_WIDTH - 1)
			player.x++;
	}
	if (IsKeyPressed(KEY_LEFT)) {
		if (player.x > 0)
			player.x--;
	}
	if (IsKeyPressed(KEY_UP)) {
		if (player.y > 0)
			player.y--;
	}
	if (IsKeyPressed(KEY_DOWN)) {
		if (player.y < MAP_HEIGHT - 1)
			player.y++;
	}

	if (IsKeyPressed(KEY_Q) && IsKeyDown(KEY_LEFT_SHIFT)) {
		CloseWindow();
	}
}

void update_3d_screen(void)
{
	ClearBackground(RAYWHITE);
	BeginMode3D(world.camera);

	for (int y = 0; y < MAP_HEIGHT; y++) {
		for (int x = 0; x < MAP_WIDTH; x++) {
			Vector3 pos = {x * TILE_SIZE, 0.0f, y * TILE_SIZE};
			DrawCube(pos, TILE_SIZE, 0.25f, TILE_SIZE, GRAY);
			DrawCubeWires(
				pos, TILE_SIZE, 0.25f, TILE_SIZE, DARKGRAY);

			Vector3 modelPos = {x * TILE_SIZE, 0.0f, y * TILE_SIZE};
			if (map[y][x] <= 94)
				DrawModel(world.sandModel, modelPos, 1.0f, WHITE);
			else if (map[y][x] >= 95)
				DrawModel(world.treeModel, modelPos, 1.0f, WHITE);
		}
	}

	Vector3 playerPos = {
		player.x * TILE_SIZE,
		0.25f + 0.25f / 2,
		player.y * TILE_SIZE
	};
	DrawModel(player.pirateModel, playerPos, 0.35f, WHITE);
	EndMode3D();
}

int main()
{
	InitWindow(1024, 768, "3D Isometric Map with Player Cube");
	SetTargetFPS(60);

	world.sandModel = LoadModel("assets/3d/beach/Sand_Flat.obj");
	world.treeModel = LoadModel("assets/3d/beach/Prop_Tree_Palm_2.obj");
	player.pirateModel = LoadModel("assets/3d/pirate.obj");
	player.texPirate = LoadTexture("assets/tile-beach-sand-sm.png");
	player.pirateModel.materials[0].maps[MATERIAL_MAP_DIFFUSE].texture = player.texPirate;
	DisableCursor();

	for (int y = 0; y < MAP_HEIGHT; y++) {
		for (int x = 0; x < MAP_WIDTH; x++) {
			map[y][x] = GetRandomValue(0, 100);
		}
	}

	while (!WindowShouldClose()) {
		UpdateCamera(&world.camera, CAMERA_FREE);
		handle_input();

		BeginDrawing();

		update_3d_screen();

		// UI Overlay
		DrawText("Use WASD to Zoom and Move Camera - Arrow Keys to "
			 "Move the Red Cube",
			 10,
			 10,
			 20,
			 DARKGRAY);

		EndDrawing();
	}

	CloseWindow();
	return 0;
}
