#include <REGX52.H>

typedef unsigned int u16;
typedef unsigned char u8;

sbit LSA = P2^2;
sbit LSB = P2^3;
sbit LSC = P2^4;

u8 code smgduan[17]=
{
0x3f,0x06,0x5b,0x4f,0x66,0x6d,0x7d,0x07,
					0x7f,0x6f,0x77,0x7c,0x39,0x5e,0x79,0x71
};//œ‘ æ0~Fµƒ÷µ
void Delay(u16 time){
	u8 i,j;
	while(time)
	{
		i=2;
		j=239;
		do
		{
			while(--j);
		}while(--i);
		time--;
	}
}

void Choose(u8 Location,Number)
{
	switch(Location)
	{
		case 1:LSC=0;LSB=0;LSA=0;break;
		case 2:LSC=0;LSB=0;LSA=1;break;
		case 3:LSC=0;LSB=1;LSA=0;break;
		case 4:LSC=0;LSB=1;LSA=1;break;
		case 5:LSC=1;LSB=0;LSA=0;break;
		case 6:LSC=1;LSB=0;LSA=1;break;
		case 7:LSC=1;LSB=1;LSA=0;break;
		case 8:LSC=1;LSB=1;LSA=1;break;
	}
	P0=smgduan[Number];
}

void main()
{
	u16 i,j;
	while(1)
	{
		for( i=0;i<8;i++){
			Choose(i+1,j);
			Delay(100);
		}
		
	};
}