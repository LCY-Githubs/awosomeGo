#include <REGX52.H>
//#include<intrins.h>	
typedef unsigned int u16;
typedef unsigned char u8;

//void Delay(u16 i);
sbit beep=P1^5;

//i=1ʱ�������ʱ10΢��
void Delay(u16 i){
	while(i--);
}

void main(){
	while(1){
		beep=~beep;
		Delay(100);
	};
}