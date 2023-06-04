#include <REGX52.H>
#include <INTRINS.H>
#include <stdio.h>

void Delay500ms()		//@12.000MHz
{
	unsigned char i, j, k;

	_nop_();
	i = 4;
	j = 205;
	k = 187;
	do
	{
		do
		{
			while (--k);
		} while (--j);
	} while (--i);
}

void Delay1ms(unsigned int xms){
		unsigned char i,j;	
		while(xms)
		{
			i = 2;
			j = 239;
				do
					{
						while (--j);
					} while (--i);
			xms--;
		}
}


void main()
{
	//printf("hello world!");
	int i = 20;
	while(1)
	{
			P2=0xFE;//1111 1110
			Delay1ms(i);	
			P2=0xFd;//1111 1101
			Delay1ms(i);	
			P2=0xFB;//1111 1011
			Delay1ms(i);	
			P2=0xF7;//1111 0111
			Delay1ms(i);	
			P2=0xEf;//1110 1111
			Delay1ms(i);	
			P2=0xDf;//1101 1101
			Delay1ms(i);	
			P2=0xBf;//1011 1101
			Delay1ms(i);	
			P2=0x7f;//0111 1101
			Delay1ms(i);	
		}
}