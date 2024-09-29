<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title}}</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css"
          integrity="sha384-xOolHFLEh07PJGoPkLv1IbcEPTNtaed2xpHsD9ESMhqIYd0nLMwNLD69Npy4HI+N" crossorigin="anonymous">
</head>
<body>
<div class="container">
    <form method="post" action="/api_demo/edit/{{.book.Id}}">
        <input type="hidden" name="_method" value="PUT" />
        <div class="form-group">
            <label for="title">Title</label>
            <input type="text" class="form-control" id="title" name="title" value="{{.book.Title}}">
        </div>
        <div class="form-group">
            <label>Content</label>
            <textarea class="form-control" name="content" rows="3">{{.book.Content}}</textarea>
        </div>
        <input type="hidden" name="id" value="{{.book.Id}}">
        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
</div>
</body>
</html>