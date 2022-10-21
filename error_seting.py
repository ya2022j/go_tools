
class ResponseError(BaseException):
    def __init__(self,msg):
        self.msg = msg
    def __str__(self):
        return self.msg




for item in range(10):
    print(item)
    if item == 8:
        raise ResponseError("ssss")
