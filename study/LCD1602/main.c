#include <REGX52.H>
#include "lcd.h"

u8 Disp[] = "Hello World";

void main(){
	u8 i;
	LcdInit();
	for(i=0;i<16;i++){
		LcdWriteData(Disp[i]);
	}
	while(1);
}