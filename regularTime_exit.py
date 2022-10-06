



import datetime
import sys
import time


# "15:32:55"
# 在特定时间停止函数的装饰器和函数


def at_sometime_exit(parameter):
    print(parameter)
    def wrapper(func):
        def inner(*args,**kwargs):
            ret = func(*args,**kwargs)
            while True:
                print("ssss")
                now_time = datetime.datetime.strptime(datetime.datetime.now().strftime("%H:%M:%S"), "%H:%M:%S")
                stop_timeline = datetime.datetime.strptime(parameter, "%H:%M:%S")
                print(now_time)
                print(stop_timeline)
                time.sleep(1)
                if now_time > stop_timeline:
                    print("ok")
                    sys.exit(0)
                return ret
        return inner
    return wrapper

def regular_time_exit(stop_time):
    while True:
        print("ssss")
        now_time = datetime.datetime.strptime(datetime.datetime.now().strftime("%H:%M:%S"),"%H:%M:%S")
        stop_timeline = datetime.datetime.strptime(stop_time, "%H:%M:%S")
        print(now_time)
        print(stop_timeline)
        time.sleep(1)
        if now_time>stop_timeline:
            print("ok")
            sys.exit(0)

@at_sometime_exit("15:55:55")
def ddd():
    print(datetime.datetime.now())
    time.sleep(1)


while True:
    ddd()





