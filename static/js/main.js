function EditFilesTable() {
  $.ajax({
    url: "/dir",
    type: "GET",
    dataType: "json",
    success: function (data) {
      var files = data;
      var table = $("#files-table");
      table.empty();
      table.append(
        "<tr><th>File</th><th>Size</th><th>Type</th><th>Actions</th></tr>"
      );
      for (var i = 0; i < files.length; i++) {
        var file = files[i];
        var row = $("<tr></tr>");
        row.append("<td>" + file.name + "</td>");
        row.append("<td>" + file.size + "</td>");
        row.append("<td>" + file.type + "</td>");
        row.append(
          '<td><button class="btn-danger" onclick="deleteFile(\'' +
            file.name +
            "')\">Delete</button></td>"
        );
        table.append(row);
      }
    },
  });
}

$(document).ready(function () {
  EditFilesTable();
});
