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
        int n;
	struct  timeval tv;
	struct  timezone tz;
        if((mkfifo(FIFO, O_CREAT) < 0) && (errno != EEXIST))
        {
                printf("不能创建FIFO\n");
                exit(1);
        }

        printf("准备读取数据\n");
        fd = open(FIFO, O_RDWR, 0);
        if(fd == -1)
        {
                perror("打开FIFO");
                exit(1);
        }

	for(int i=0; i<100; i++)
        {
                if((n = write(fd, &i, sizeof(int))) == -1)
                {
                        if(errno == EAGAIN) printf("没有数据\n");
                }
		else if(n != sizeof(int))
		{
			printf("client: write %d byte\n",n);
			close(fd);
			exit(0);
		}

                gettimeofday(&tv,&tz);
                printf("从FIFO读取的数据为：%dth, %lds,%ldus\n", i, tv.tv_sec, tv.tv_usec);
		sleep(1);
        }
	close(fd);
}
