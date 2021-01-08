Simple web-app for p4p UFC top fighter rankings.<br>
The target is to learn Redis with concurrently safe Redis connections pool on a simple web-server.
Build with **redis** using <a href="https://github.com/gomodule/redigo">redigo</a> library
<br>
Data is just an up-to-date lorem from <a href="https://www.ufc.com/rankings">ufc website</a>

<table>
<tr>
<th>Method</th>
<th>Path</th>
<th>Function</th>
</tr>
  <tr>
<td>GET</td>
    <td><b>/fighter?id=1</b></td>
    <td>Show details on specific fighter (id provided in query string)</td>
  </tr>
<tr>
<td>
POST
</td>
<td><b>/win || /lose</b></td>
<td>Add win to a specific fighter (accordint to id in request body)</td>
</tr>

<tr>
<td>POST</td>
<td><b>/lose</b></td>
<td>Add lose to a specific fighter (accordint to id in request body)</td>
</tr>
</table>

>EXAMPLE OF USE:
>> curl -i -L -d "id=3" localhost:8080/lose
<br>
>> curl -i -L -d "id=3" localhost:8080/win

>-d, --data <data> Send specified data in POST request. Details provided below.<br>
>-i, --include Include HTTP headers in the output.<br>
>-L, --location Follow redirects.<br>


PS: Uncomment "ImportData()" function to fill the DB with json data. 
