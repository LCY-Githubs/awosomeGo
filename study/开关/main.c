#include <REGX52.H>


unsigned int i=0;
void main(){
	char arr[]={
		0xfe,//1111 1110
		0xfd,//1111 1101
		0xfb,
		0xf7,
		0xef,
		0xdf,
		0xbf,
		0x7f
	};
	while(1){
		if(P3_1==0){
			if(i==8){
				i=0;
			}
			P2=arr[i++];
		}else{
			
		}
	}
}