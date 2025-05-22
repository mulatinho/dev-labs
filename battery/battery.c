#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <dirent.h>
#include <limits.h>
#include <sys/stat.h>

#define M_BAT_BOLD   "\033[1m"
#define M_BAT_CLEAR  "\033[m"
#define M_BAT_RED    "\033[41;5m"
#define M_BAT_YELLOW "\033[43;5m"
#define M_BAT_GREEN  "\033[42;5m"
#define M_BAT_SCREEN "\033[2J\033[H"

#define M_BAT_SEP    "|"
#define M_BAT_LINE   "  -  -  -  -  -  - "
#define M_BAT_INIT   " _________________ "
#define M_BAT_END    " ''''''''''''''''' "
#define M_BAT_BLANK  "                   "

const char *read_data(char *filename)
{
	FILE *fp;
	char buffer[NAME_MAX] = {0};
	struct stat st;

	if (stat(filename, &st) == -1)
		return NULL;

	if (!(fp = fopen(filename, "r")))
		return NULL;

	if (!fgets(buffer, sizeof(buffer) - 1, fp))
		return NULL;

	fclose(fp);

	return strdup(buffer);
}

static void update_screen(const char *model_name, const int capacity)
{

	int level = 100;

	printf(M_BAT_SCREEN 
	  M_BAT_BOLD  " Battery Monitor " M_BAT_CLEAR "\n"
	  M_BAT_BOLD  " Model: " M_BAT_CLEAR "%s" M_BAT_CLEAR
	  M_BAT_BOLD  " Current Capacity: %2d%" M_BAT_CLEAR "\n",
	  model_name, capacity);

	printf(M_BAT_CLEAR M_BAT_INIT  M_BAT_CLEAR "\n");

	for (; level >= 0; level -= 5) {
		int print_color = 0;

		if (capacity >= level)
			print_color++;

		if (level >= 70 && print_color)
	  		printf(M_BAT_CLEAR M_BAT_SEP M_BAT_GREEN  M_BAT_LINE M_BAT_CLEAR "|\n");
		else if (level >= 30 && level < 70 && print_color)
	  		printf(M_BAT_CLEAR M_BAT_SEP M_BAT_YELLOW M_BAT_LINE M_BAT_CLEAR "|\n");
		else if (level < 30 && print_color)
	  		printf(M_BAT_CLEAR M_BAT_SEP M_BAT_RED    M_BAT_LINE M_BAT_CLEAR "|\n");
		else
	      	printf(M_BAT_CLEAR M_BAT_SEP M_BAT_CLEAR  M_BAT_BLANK M_BAT_CLEAR "|\n");
	}

	printf(M_BAT_CLEAR M_BAT_END M_BAT_CLEAR "\n");
}

int main(int argc, char **argv)
{
	const char *capacity_data = read_data("/sys/class/power_supply/BAT0/capacity");
	const char *model_name = read_data("/sys/class/power_supply/BAT0/model_name");
	int capacity = atoi(capacity_data);

	update_screen(model_name, capacity);
	return 0;
}
