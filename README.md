# BE-Basic

回答 [问题](https://docs.google.com/document/d/10GyxdqfAiIamEjlfwkx0q1cvlVZ1FE9kGVdSbWR0Wqs/edit)
并按照下面 API 说明提交答案。

请求方式: POST  
地址：http://bt1.cool2645.cc/quiz  
参数：name  
说明：注册，注册后服务端会返回一个 token。  

请求方式: PUT  
地址：http://bt1.cool2645.cc/quiz  
参数：token, name  
说明：可以用来修改姓名。  

请求方式: DELETE  
地址：http://bt1.cool2645.cc/quiz  
参数：token  
说明：提交试卷。提交后你的名字将显示在首页上。请注意，提交后答案将不可修改！  

请求方式: GET  
地址：http://bt1.cool2645.cc  
说明：显示所有已经提交试卷的人的姓名。  

请求方式: POST  
地址：http://bt1.cool2645.cc  
参数：pwd  
说明：显示所有姓名及答案，需要提供管理员密码。  

请求方式: POST  
地址：http://bt1.cool2645.cc/question/{id}  
路由参数：id：题号  
参数：token, answer  
说明：上传回答。  

请求方式: PUT  
地址：http://bt1.cool2645.cc/question/{id}  
路由参数：id：题号  
参数：token, answer  
说明：修改答案。  

请求方式: GET  
地址：http://bt1.cool2645.cc/question/{id}  
路由参数：id：题号  
参数：token  
说明：查看答案。  

