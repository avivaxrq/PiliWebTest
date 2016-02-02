$(function() {


    function downloadFile(domain, key, fileSize) {
          $.ajax({
            url: UpHost,
            data: data,
            cache: false,
            contentType: false,
            processData: false,
            type: 'POST'
        }).fail(function(data) {

            content += "<tr><td><%- UpHost %></td>";
            content += "<td colspan=4 class='fail'>失败：<%- newStatus %>&nbsp;&nbsp;&nbsp;<%- newRText %></td></tr>";

            var template = _.template(content);
            content = template({
                UpHost: UpHost,
                downloadSize: downloadSize,
                newStatus: data.status,
                newRText: data.responseText
            });

            $(content).appendTo('.js-upload-area');

            downloadFile(domain, 'helloworld.txt?_' + (+new Date()), defaultFileSize);
        }).pipe(function(data) {
            var now = +new Date();
            var time = (now - t) / 1000;
            var speed = (size / 1024 / time).toFixed(2);

            content += "<tr><td><%- UpHost %></td>";
            content += "<td><%- downloadSize %>Kb</td>";
            content += "<td><%- time %>s</td>";
            content += "<td><%- speed %>Kb/s</td>";
            content += "<td>成功</td></tr>";

            var template = _.template(content);
            content = template({
                UpHost: UpHost,
                downloadSize: downloadSize,
                time: time,
                speed: speed
            });

            $(content).appendTo('.js-upload-area');

            downloadFile(domain, data.key, size);
        });
    }
    function downloadFile(domain, key, fileSize) {
        var t = +new Date();
        var downloadInnerHtml = '';
        var url = domain + '/' + key;
        var downloadSize = (fileSize / 1024).toFixed(2)

        $.ajax({
            type: 'GET',
            url: url
        }).success(function(data) {
            var now = +new Date();
            var time = (now - t) / 1000;
            var speed = (downloadSize / time).toFixed(2);
            downloadInnerHtml += "<tr><td><%- domain %></td>";
            downloadInnerHtml += "<td><%- downloadSize %>Kb</td>";
            downloadInnerHtml += "<td><%- time %>s</td>";
            downloadInnerHtml += "<td><%- speed %>Kb/s</td>";
            downloadInnerHtml += "<td>成功</td></tr>";

            var template = _.template(downloadInnerHtml);
            downloadInnerHtml = template({
                domain: domain,
                url: url,
                downloadSize: downloadSize,
                time: time,
                speed: speed
            });

            $(downloadInnerHtml).appendTo('.js-download-area');

        }).fail(function() {
            var now = +new Date();
            var time = (now - t) / 1000;

            downloadInnerHtml += "<tr><td><%- domain %></td>";
            downloadInnerHtml += "<td><%- downloadSize %>s</td>";
            downloadInnerHtml += "<td><%- time %>s</td>";
            downloadInnerHtml += "<td>-</td>";
            downloadInnerHtml += "<td>失败</td></tr>";

            var template = _.template(downloadInnerHtml);
            downloadInnerHtml = template({
                domain: domain,
                url: url,
                downloadSize: downloadSize,
                time: time
            });

            $(downloadInnerHtml).appendTo('.js-download-area');
        });
    }

    try {
        var up_host = $('#up_host').val();
        var upload_host = $('#upload_host').val();
        uploadFile(up_host);
        uploadFile(upload_host);
    } catch (e) {
        $('<tr><td colspan=5>浏览器版本太低，不支持Blob对象，无法自动创建文件上传至七牛。</td></tr>').appendTo('.js-upload-area');

        downloadFile(domain, 'helloworld.txt?_' + (+new Date()), defaultFileSize);
    }

    $('.js-time-area').text(getCurrentTime());
});
