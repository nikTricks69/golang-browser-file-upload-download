<html>

<head>
 <meta http-equiv="cache-control" content="no-cache, must-revalidate, post-check=0, pre-check=0" />
  <meta http-equiv="cache-control" content="max-age=0" />
  <meta http-equiv="expires" content="0" />
  <meta http-equiv="expires" content="Tue, 01 Jan 1980 1:00:00 GMT" />
  <meta http-equiv="pragma" content="no-cache" />
	<title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
	<input type="file" name="uploadFile" />
	<input type="submit" value="upload" />
</form>
<button onclick="location.href = '/files';" id="myButton" class="float-left submit-button" >View files</button>
</body>
</html>