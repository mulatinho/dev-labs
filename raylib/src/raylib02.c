#include "raylib.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define MLT_KEYBOARD_SIZE 32
#define MLT_KEYBOARD_SIZE_ELEMENTS 7
#define MLT_KEYBOARD_SIZE_WIDTH (MLT_KEYBOARD_SIZE * MLT_KEYBOARD_SIZE_ELEMENTS)

typedef struct {
	int screen_width;
	int screen_height;
	Texture2D tileset_keyboard;
} mlt_window_t;

mlt_window_t win;

void handle_input(void) 
{
	if (IsKeyPressed(KEY_Q)) {
		if (IsTextureValid(win.tileset_keyboard))
			UnloadTexture(win.tileset_keyboard);
		CloseWindow();
		exit(EXIT_SUCCESS);
	}
}

void update_vector_xy(int *x, int *y)
{
	*x = (*x + MLT_KEYBOARD_SIZE) % 640;
	*y = 64;
}

void update_vector_txy(int *x, int *y)
{
	*x = *x / MLT_KEYBOARD_SIZE == MLT_KEYBOARD_SIZE_ELEMENTS ? 0 : *x + MLT_KEYBOARD_SIZE;
	if (*x >= MLT_KEYBOARD_SIZE_WIDTH)
		*y += MLT_KEYBOARD_SIZE;

	if (*y >= (MLT_KEYBOARD_SIZE * 2))
		*y = 0;

	fprintf(stdout, "%s:%s - x:%3d, y:%3d \n", __FILE__, __func__, *x, *y);
}

int main(void)
{
	win.screen_width = 800;
	win.screen_height = 600;

	InitWindow(win.screen_width, win.screen_height, "Raylib Keyboard");
	win.tileset_keyboard = LoadTexture("assets/tile_keyboard.png");

	SetTargetFPS(60);

	int xt = 0, yt = 0, x = 0, y = 0;
	while (!WindowShouldClose()) {
		handle_input();
		BeginDrawing();
		ClearBackground(WHITE);

		Rectangle src = {xt, yt, MLT_KEYBOARD_SIZE, MLT_KEYBOARD_SIZE};
		Vector2 dest = {x, y};

		DrawTextureRec(win.tileset_keyboard, src, dest, WHITE);

		EndDrawing();
		update_vector_xy(&x, &y);
		update_vector_txy(&xt, &yt);
		usleep(150 * 1000);
	}

	return 0;
}
