$(function() {
    var domain = $('#domain').val();
    var defaultFileSize = 3379; //byte

    function uniqueFileName() {
        var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXTZabcdefghiklmnopqrstuvwxyz";
        var string_length = 15;
        var randomstring = '';
        for (var i = 0; i < string_length; i++) {
            var rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum, rnum + 1);
        }

        var t = +new Date();
        return randomstring + t + '.txt';
    }

    function getCurrentTime() {
        var date = new Date();
        return date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate() + '  ' + date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds();
    }

    function uploadFile(UpHost) {

        var fileContent = "";
        fileContent += "     IO                                                              N7O\r\n";
        fileContent += "     7IZD                                                          NO$7Z\r\n";
        fileContent += "     77IIIO                                                      N7III7O\r\n";
        fileContent += "     777IZ8Z7ON                MD8OOZZZZOO8NM                 D7$OOII77N\r\n";
        fileContent += "     Z77$IIIII7NDZ$ZZOO8OO$777IIIIIIIIIIIIIIII777$ZO88OOZ$$8N77IIII$77O\r\n";
        fileContent += "      Z7777II$N7IIIIIII777IIIIIIIIIIIIIIIIIIIIIIII7777IIIIII7$D?II7777\r\n";
        fileContent += "       D$777ZZIIIIII777IIIIIIIIIIIIIIIIIIIIIIIIIIIIII7777IIIII777777$\r\n";
        fileContent += "         $7777777777IIIIIIIIIIIIII7IIIIIIIIIIIIIIIIIIIII7777777777Z\r\n";
        fileContent += "           D7$7777IIIIIIIIIIIII7MMZIIII7NMD7IIIIIIIIIIIII777777$O\r\n";
        fileContent += "             MN$7IIIIIIIIIIIIIIZMMNIIIIZMMM7IIIIIIIIIIIIII77$DM\r\n";
        fileContent += "            M8I7IIIIIIIIIIIIIIII77IIIIIII$7IIIIIIIIIIIIIIII77?DM\r\n";
        fileContent += "          MM~77IIIIIIIII7II77I???+++===++++??IIII7IIIIIIIII7777ZMM\r\n";
        fileContent += "        MM?.777II7I+~,................................,~+II7777+.IMM\r\n";
        fileContent += "     MMM+.,77$7,.....                                     ..=I7$7..OMM\r\n";
        fileContent += "   MZ=..,I7$:. .....   ....                     ....       ..,,=77=...?OM\r\n";
        fileContent += "   M,.=II7=,......    .:~~~:,                 ,~~~~:.       .,,,,:III=..M\r\n";
        fileContent += "   MZ.,?I:,,.....    .:~~~~~~:.             ,~~~~~~~~.       .,,,,=I?,.OM\r\n";
        fileContent += "    M,..:,,,....      :~~~~~~~.            .~~~~~~~~,         ,,,,:~..$M\r\n";
        fileContent += "     MMM~,,....        .,,:::,              .,::::,.        ..,,,,,?MM\r\n";
        fileContent += "       M:,,....                                             ..,,,,,?M\r\n";
        fileContent += "       M~,,,....                                             .,,,,,IM\r\n";
        fileContent += "       M7,,,. ..                                           ..,,,,,,M\r\n";
        fileContent += "        D,,,,. ...                                         .,,,,,,?M\r\n";
        fileContent += "         M+,,,,....                                       .,,,,,,IM\r\n";
        fileContent += "          M8,,,,,...                                    .,,,,,,?M\r\n";
        fileContent += "            DI~,,,,...                               .,,,,,,,=7\r\n";
        fileContent += "               MN$~,,,,,,,,.                    ,,,,,,,,,+ZM\r\n";
        fileContent += "                   NMNZ?=~,,,,,,,,,,,,,,,,,,,,,,,,~=IOMM\r\n";
        fileContent += "                         MMDZI7777777777$777777DMM\r\n";
        fileContent += "                       MD$IIIIIIIIIIIIIIIIIIIIIIII8M\r\n";
        fileContent += "                     N$7I7IIIIIIIIII?IIIIIIIIIIIIII77N\r\n";
        fileContent += "                   OIIII7IIIIIIIII:....?IIIIIIIIIIIII77Z\r\n";
        fileContent += "                 N7IIII7IIIIIIIII::~..:.IIIIIIIIII7IIII7$D\r\n";
        fileContent += "                ZIIIII7IIIIIIIIII~...~,.7IIIIIIIII77IIII77$\r\n";
        fileContent += "               D7IIII77IIIIIIIIIII+....7IIIIIIIIIII77IIII77D\r\n";
        fileContent += "               MM87777IIIIIIIIIIIIIIIIIIIIIIIIIIIII77$II7OM\r\n";
        fileContent += "                     DIIIIIII777IIIIIIIIII777IIIIII77D\r\n";
        fileContent += "                     8IIIIIIIIIII7777777IIIIIIIIIII778\r\n";
        fileContent += "                     DIIIIIII7777777777777IIIIIIII777D\r\n";
        fileContent += "                      III77$$$$$Z8       8Z$$$777I$77\r\n";
        fileContent += "                       D8OZOONM             MN8OOO8D\r\n\r\n";
        fileContent += "  文件生成时间：" + getCurrentTime();


        var blob = new Blob([fileContent], {
            'type': "text/plain"
        });
        var size = blob.size;
        var token = $('#token').val();
        var downloadSize = (size / 1024).toFixed(2);


        var data = new FormData();
        data.append('token', token);
        data.append('file', blob);
        data.append('key', uniqueFileName());

        var t = +new Date();
        var content = '';

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
