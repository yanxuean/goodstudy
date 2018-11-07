#include<errno.h>
#include<sys/stat.h>
#include<fcntl.h>
#include<sys/time.h>
#include "stdio.h"
#include <stdlib.h>
#include <unistd.h>

//int gettimeofday(struct  timeval*tv,struct  timezone *tz )

#define FIFO "/tmp/my_fifo"
//本程序从一个FIFO读数据，并把读到的数据打印到标准输出
//如果读到字符“Q”，则退出
int main(int argc, char** argv)
{
        char buf_r[100];
        int fd;
        int nread;
	struct  timeval tv;
        struct  timezone tz;

        if((mkfifo(FIFO, O_CREAT) < 0) && (errno != EEXIST))
        {
                printf("不能创建FIFO\n");
                exit(1);
        }

        fd = open(FIFO, O_RDONLY, 0);
        if(fd == -1)
        {
                perror("打开FIFO");
                exit(1);
        }
        
        printf("准备读取数据\n");
        for(int i=0; ;i++)
	{
                if((nread = read(fd, buf_r, sizeof(int))) == -1)
                {
                        if(errno == EAGAIN) printf("没有数据\n");
                }
		else if(nread == 0)
		{
			printf("server: recv eof\n");
			close(fd);
			exit(0);
		}

                //假设取到Q的时候退出
                //if(buf_r[0]=='Q') break;

                //buf_r[nread]=0;
                printf("从FIFO读取的数据为：%d\n", *(int *)buf_r);
                //sleep(1);
		gettimeofday(&tv,&tz);
                printf("从FIFO读取的数据为：%dth, %lds,%ldus\n", i, tv.tv_sec, tv.tv_usec);
        }

}
