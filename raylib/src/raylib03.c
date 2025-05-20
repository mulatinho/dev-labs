#include "raylib.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define MLT_WINDOW_WIDTH 800
#define MLT_WINDOW_HEIGHT 600

#define MLT_KEYBOARD_SIZE 32
#define MLT_KEYBOARD_SIZE_ELEMENTS 7
#define MLT_KEYBOARD_SIZE_WIDTH (MLT_KEYBOARD_SIZE * MLT_KEYBOARD_SIZE_ELEMENTS)

typedef enum {
	MLT_KEY_1, MLT_KEY_2, MLT_KEY_3, MLT_KEY_4, MLT_KEY_5, MLT_KEY_6,
	MLT_KEY_7, MLT_KEY_8, MLT_KEY_9, MLT_KEY_A, MLT_KEY_B, MLT_KEY_C,
	MLT_KEY_D, MLT_KEY_E, MLT_KEY_F, MLT_KEY_G, MLT_KEY_H, MLT_KEY_I,
	MLT_KEY_J, MLT_KEY_K, MLT_KEY_L, MLT_KEY_M, MLT_KEY_N, MLT_KEY_O,
	MLT_KEY_P, MLT_KEY_Q, MLT_KEY_R, MLT_KEY_S, MLT_KEY_T, MLT_KEY_U,
	MLT_KEY_V, MLT_KEY_W, MLT_KEY_X, MLT_KEY_Y, MLT_KEY_Z
} mlt_key_t;

typedef struct {
	int posx;
	int posy;
	Texture2D tileset_keyboard;
} mlt_controls_t;

typedef struct {
	int screen_width;
	int screen_height;
	Texture2D tileset_jungle;
	Texture2D tileset_keyboard;
} mlt_window_t;

mlt_window_t win;

void mlt_handle_input(void) 
{
	if (IsKeyPressed(KEY_Q)) {
		if (IsTextureValid(win.tileset_jungle))
			UnloadTexture(win.tileset_jungle);
		if (IsTextureValid(win.tileset_keyboard))
			UnloadTexture(win.tileset_keyboard);
		CloseWindow();
		exit(EXIT_SUCCESS);
	}
}

void mlt_texture_draw(int x, int y, mlt_key_t key)
{
	int xt = 0, yt = 0;

	switch (key) {
		case MLT_KEY_1: xt = 32*0; yt = 32*0; break;
		case MLT_KEY_2: xt = 32*1; yt = 32*0; break;
		case MLT_KEY_3: xt = 32*2; yt = 32*0; break;
		case MLT_KEY_4: xt = 32*3; yt = 32*0; break;
		case MLT_KEY_5: xt = 32*4; yt = 32*0; break;
		case MLT_KEY_6: xt = 32*5; yt = 32*0; break;
		case MLT_KEY_7: xt = 32*6; yt = 32*0; break;
		case MLT_KEY_8: xt = 32*0; yt = 32*1; break;
		case MLT_KEY_9: xt = 32*1; yt = 32*1; break;
		case MLT_KEY_A: xt = 32*2; yt = 32*1; break;
		case MLT_KEY_B: xt = 32*3; yt = 32*1; break;
		case MLT_KEY_C: xt = 32*4; yt = 32*1; break;
		case MLT_KEY_D: xt = 32*5; yt = 32*1; break;
		case MLT_KEY_E: xt = 32*6; yt = 32*1; break;
		case MLT_KEY_F: xt = 32*0; yt = 32*2; break;
		case MLT_KEY_G: xt = 32*1; yt = 32*2; break;
		case MLT_KEY_H: xt = 32*2; yt = 32*2; break;
		case MLT_KEY_I: xt = 32*3; yt = 32*2; break;
		case MLT_KEY_J: xt = 32*4; yt = 32*2; break;
		case MLT_KEY_K: xt = 32*5; yt = 32*2; break;
		case MLT_KEY_L: xt = 32*6; yt = 32*2; break;
		case MLT_KEY_M: xt = 32*0; yt = 32*3; break;
		case MLT_KEY_N: xt = 32*1; yt = 32*3; break;
		case MLT_KEY_O: xt = 32*2; yt = 32*3; break;
		case MLT_KEY_P: xt = 32*3; yt = 32*3; break;
		case MLT_KEY_Q: xt = 32*4; yt = 32*3; break;
		case MLT_KEY_R: xt = 32*5; yt = 32*3; break;
		case MLT_KEY_S: xt = 32*6; yt = 32*3; break;
		case MLT_KEY_T: xt = 32*0; yt = 32*4; break;
		case MLT_KEY_U: xt = 32*1; yt = 32*4; break;
		case MLT_KEY_V: xt = 32*2; yt = 32*4; break;
		case MLT_KEY_W: xt = 32*3; yt = 32*4; break;
		case MLT_KEY_X: xt = 32*4; yt = 32*4; break;
		case MLT_KEY_Y: xt = 32*5; yt = 32*4; break;
		case MLT_KEY_Z: xt = 32*6; yt = 32*4; break;
		default: break;
	}

	Rectangle src = {xt, yt, MLT_KEYBOARD_SIZE, MLT_KEYBOARD_SIZE};
	Vector2 dest = {x, y};

	DrawTextureRec(win.tileset_keyboard, src, dest, WHITE);
}

int mlt_create_controls(void)
{
	int column = 64, line = 64;

	mlt_texture_draw(line, column, MLT_KEY_1); column += 32;
	mlt_texture_draw(line, column, MLT_KEY_9); column += 32;
	mlt_texture_draw(line, column, MLT_KEY_8); column += 32;
	mlt_texture_draw(line, column, MLT_KEY_5); column += 32;

	line += 32; column = 64;

	mlt_texture_draw(line, column, MLT_KEY_A); column += 32;
	mlt_texture_draw(line, column, MLT_KEY_L); column += 32;
	mlt_texture_draw(line, column, MLT_KEY_E); column += 32;
	mlt_texture_draw(line, column, MLT_KEY_X); column += 32;
}

int main(void)
{
	win.screen_width  = MLT_WINDOW_WIDTH;
	win.screen_height = MLT_WINDOW_HEIGHT;

	InitWindow(win.screen_width, win.screen_height, "Raylib Keyboard");
	win.tileset_keyboard = LoadTexture("assets/tile_keyboard.png");

	SetTargetFPS(60);

	while (!WindowShouldClose()) {
		mlt_handle_input();
		BeginDrawing();
		ClearBackground(BLANK);

		mlt_create_controls();

		EndDrawing();
		usleep(150 * 1000);
	}

	return 0;
}
