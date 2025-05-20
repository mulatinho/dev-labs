#include "raylib.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

typedef struct {
	int screenWidth;
	int screenHeight;
	Texture2D tileset_jungle;
	Texture2D tileset_keyboard;
} mlt_window_t;

mlt_window_t win;

void handle_input(void) 
{
	if (IsKeyPressed(KEY_Q)) {
		UnloadTexture(win.tileset);
		CloseWindow();
		exit(EXIT_SUCCESS);
	} else
		fprintf(stdout, "q is not pressed\n");
}

void update_vector_xy(int *x, int *y)
{
	*x = (*x + 64) % 640;
	*y = (*y + 64) % 512;
}

void update_vector_txy(int *x, int *y)
{
	*x = *x % 196 ? 0 : *x + 64;
	*y = *y % 128 ? 0 : *y + 64; 
}
int main(void)
{
	win.screenWidth = 800;
	win.screenHeight = 600;

	InitWindow(win.screenWidth, win.screenHeight, "Raylib Tileset Example");
	win.tileset_jungle = LoadTexture("assets/tile_jungle_plants_objects.png");


	SetTargetFPS(60);

	int xt = 0, yt = 0, x = 0, y = 0;
	while (!WindowShouldClose()) {
		handle_input();
		BeginDrawing();
		ClearBackground(BLANK);

		Rectangle src = {xt, yt, 64, 64};
		Vector2 dest = {x, y};

		DrawTextureRec(win.tileset_jungle, src, dest, WHITE);

		EndDrawing();
		update_vector_xy(&x, &y);
		update_vector_txy(&xt, &yt);
		fprintf(stdout, "xt: %d, yt: %d\n", xt, yt);
		usleep(150 * 1000);
	}

	return 0;
}
