
        let url = window.location.origin.replace("http","ws")
        let socket = new WebSocket(url+"/select")
        
        console.log("請輸入要查詢的檔案名稱");

        socket.onopen = function(){
            console.log("連線到:" + socket);
        }

        socket.close = function(e){
            console.log("連線失敗 (" + "錯誤代碼:"+ e.code + ")");
        }

        socket.onmessage = function(e){
            console.log("Server回傳:"+e.data)
            var a = "data:image/jpeg;base64,"
            var b = a + e.data
            document.getElementById('img').src = b
            var img = document.getElementById('img');
            img.style.width = 400;
        }

        socket.onerror = function(e){
            console.log("Server錯誤:" + e.data);
        }

        // $("button[class='del']").click(function(){
        //     var delId = $(this).attr('id');
        //     var name = $(this).attr('name');
        //     $.ajax({
        //         type:'POST',
        //         url:'${"127.0.0.1:8080/del"}',
        //         data:{'Id':delId},
        //         success:function(data){
        //             alert("成功刪除:<"+name+">")
        //             window.location.reload();
        //         },
        //         error:function(data){
        //             alert("刪除失敗")
        //         }
        //     })
        // })
        
        function select(){  
            var msg2 = document.getElementById('saveName').value;
            socket.send(msg2)
        };
