#include <stdlib.h>
#include <stdio.h>
#include <jpeglib.h>

#define ALIGN_SIZE 16



int main(){
	printf("hello turbo\n");
	//printf("%d", JDCT_ISLOW);
}

int read_jpeg_file(char * filename){
	struct jpeg_decompress_struct cinfo;

// 	// struct my_error_mgr jerr;


	FILE * infile;
// 	// JSAMPARRY buffer;
	int row_stride;

	if ((infile = fopen(filename, "rb")) == NULL) {
		fprintf(stderr, "can't open %s\n", filename);
		return 0;
	}

// 	// cinfo.err = jpeg_std_error(&jerr.pub);
// 	// jerr.pub.error_exit = my_error_exit;
// 	// jpeg_destroy_decompress(&cinfo);
// 	// fclose(infile);
	jpeg_create_decompress(&cinfo);

	return 0;
}