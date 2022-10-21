


 wget   https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-x86_64.tar.bz2;
 # 安装bzip2

 yum -y install bzip2 ;
  tar -jxvf phantomjs-2.1.1-linux-x86_64.tar.bz2;
  mv phantomjs-2.1.1-linux-x86_64 /usr/local/src/phantomjs;
  ln -sf /usr/local/src/phantomjs/bin/phantomjs /usr/local/bin/phantomjs;
  yum -y  install fontconfig freetype2;
  #测试版本号
  phantomjs -v;


#
#
# 新版的 selenium已经放弃PhantomJS,所以才会出现module ‘selenium.webdriver’ has no attribute ‘PhantomJS’。PhantomJS安装都是最新版本，pip install selenium安装的也是最新版本。
# pip3 uninstall selenium
#  pip3 install -U selenium==3.3.0
#from selenium import webdriver
#
#driver=webdriver.PhantomJS(executable_path='/usr/local/bin/phantomjs')
#url = "https://www.baidu.com/"
#driver.get(url)
## 最大化浏览器
#driver.maximize_window()
#
## 截取登录框的页面保存到相应位置
#driver.save_screenshot("baidu.png")
## 模拟点击按钮跳转体育页面
## 退出驱动关闭所有窗口
#driver.quit()
