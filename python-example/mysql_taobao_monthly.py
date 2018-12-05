#coding:utf-8
#通用抓取脚本爬虫
#抓取从主题列表分页，内容列表分页中的图片
import urllib
import requests
from urlparse import urlsplit
import urllib2
import os
import sys


hrefLeft='http://mysql.taobao.org'	#主题前缀 默认为空
titleFlge='a target="_top" class="main" '	#主题标志
titleS='href="'	#主题链接开始
titleE='">'	#主题链接结束
PageTitle=20	#每页主题数量
DetailSiteSuffix=''	#内容连接类型  .html  .shtml
headers={'User-Agent':'Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6'}

pageFlge='<small class="datetime muted"' #主题前缀

pageTitleS='<a href="' #主题标题开始
pageTitleE='</a>' #主题标题结束




def OpenPageSite(url):
	global hrefLeft
	global titleS
	global titleFlge
	global PageTitle
	global DetailSiteSuffix
	global headers


	req = urllib2.Request(url=url,headers=headers)
	contentMap = urllib2.urlopen(req).read()
	title=contentMap.find(titleFlge)
	href=contentMap.find(titleS,title)
	html=contentMap.find(titleE,href)

	print(len(contentMap),title,href,html)
	if (title==0 and href==0 and html==0)or(title==- 1 and href==- 1 and html==- 1):
		print('     				Detail not find')

	i=0
	PageTitle=contentMap.count(titleFlge)
	url=['']*PageTitle
	while title!=- 1 and href!=- 1 and html!=- 1 and i<len(url):
		webAdd=hrefLeft+contentMap[href+len(titleS):html]+DetailSiteSuffix
		url[i]=webAdd
		title=contentMap.find(titleFlge,html)
		href=contentMap.find(titleS,title)
		html=contentMap.find(titleE,href)
		i=i+ 1
	else:
		print('			page find end')

	i=0
	while i<len(url):
		OpenPage(url[i])
		i=i+1

def OpenPage(url):
	global headers
	global pageFlge
	global pageTitleS
	global pageTitleE

	f = open("out.txt",'a+')

	if(len(url)==0):
			return
	print(url)

	reload(sys)
	sys.setdefaultencoding('utf-8')
	type = sys.getfilesystemencoding()

	req = urllib2.Request(url=url,headers=headers)
	contentMap = urllib2.urlopen(req).read().decode('utf-8').encode(type)
	title=contentMap.find(pageFlge)
	href=contentMap.find(pageTitleS,title)
	html=contentMap.find(pageTitleE,href)

	if (title==0 and href==0 and html==0)or(title==- 1 and href==- 1 and html==- 1):
		print('     				Detail not find')

	i=0
	PageTitle=contentMap.count(pageFlge)
	url=['']*PageTitle
	while title!=- 1 and href!=- 1 and html!=- 1 and i<len(url):

		_url= contentMap[href+len(pageTitleS):html].strip()
		#print hrefLeft
		#print _url
		print >> f, hrefLeft+_url.replace("\n","").replace('target="_blank">','').replace('"','')

		title=contentMap.find(pageFlge,html)
		href=contentMap.find(pageTitleS,title)
		html=contentMap.find(pageTitleE,href)
		i=i+ 1
	else:
		print('			page find end')


OpenPageSite('http://mysql.taobao.org/monthly');
