<!DOCTYPE html>
<html>
   <head>
      <title>Login-User</title>
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
     <!-- 新 Bootstrap 核心 CSS 文件 -->
     <link href="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">

     <!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
     <script src="https://cdn.staticfile.org/jquery/2.1.1/jquery.min.js"></script>

     <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
     <script src="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
   </head>
   <body>

        <div style="width:40%; margin:10% auto">
          <h3 style="text-align: center;">User Login</h3>
          <br>
          <br>
          <form class="form-horizontal" role="form">
            <div class="form-group">
              <label for="firstname" class="col-sm-2 control-label">用户名</label>
              <div class="col-sm-10">
                <input type="text" class="form-control" id="firstname" placeholder="请输入用户名  ">
              </div>
            </div>
            <div class="form-group">
                <label for="inputPassword" class="col-sm-2 control-label">密码</label>
                <div class="col-sm-10">
                  <input type="password" class="form-control" id="inputPassword" placeholder="请输入密码">
                </div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-2 col-sm-10">
                <div class="checkbox">
                  <label>
                    <input type="checkbox">请记住我
                  </label>
                </div>
              </div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-2 col-sm-10">
                <button type="submit" class="btn btn-default">登录</button>
              </div>
            </div>
          </form>
        </div>
   </body>
</html>