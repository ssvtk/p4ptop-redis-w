Simple web-app for p4p UFC top fighter rankings.<br>
TTarget is to learn redis with concurrently safe Redis connections pool, using simple web-server.
Build with **redis** using <a href="https://github.com/gomodule/redigo">redigo</a> library
<br>
Data is just an up-to-date lorem from <a href="https://www.ufc.com/rankings">ufc website</a>

<table>API
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
<td><b>/win</b></td>
<td>Add win to a specific fighter (accordint to id in request body)</td>
</tr>

<tr>
<td>GET</td>
<td><b>/top</b></td>
<td>List of top 3 according to wins</td>
</tr>
</table>