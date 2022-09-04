<html>
<head>
	<title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://localhost:8080/upload" method="post">
	<input type="file" name="uploadFile" />
	<input type="submit" value="upload" />
</form>
<button onclick="location.href = '/files';" id="myButton" class="float-left submit-button" >View files</button>
</body>
</html>