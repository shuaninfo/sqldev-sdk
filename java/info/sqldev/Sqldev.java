package info.sqldev;

import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.Reader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.net.URLEncoder;
import java.util.HashMap;
import java.util.Map;

import info.sqldev.user.*;
import info.sqldev.instance.*;
import info.sqldev.sqlquery.*;
import info.sqldev.util.JsonUtil;

import static info.sqldev.sign.Signer.getSignature;

public class Sqldev {
    private String endpoint;
    private String appKey;
    private String appSecret;

    public Sqldev(String endpoint, String appKey, String appSecret) {
        this.endpoint = endpoint;
        this.appKey = appKey;
        this.appSecret = appSecret;
    }

    private byte[] sendRequest(String method, String path, Map<String, String> payload) throws Exception {
        Map<String, String> params = new HashMap<String, String>();
        for (Map.Entry<String, String> entry : payload.entrySet()) {
            params.put(entry.getKey(), entry.getValue());
        }
        params.put("app_key", appKey);
        params.put("current_time", String.valueOf(System.currentTimeMillis()));

        String sign = getSignature(params, appSecret);
        params.put("sign", sign);

        URL url = new URL(endpoint + path);
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        conn.setRequestMethod(method);
        conn.setDoOutput(true);
        conn.setDoInput(true);
        conn.setUseCaches(false);
        conn.setInstanceFollowRedirects(true);
        conn.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");
        conn.connect();

        StringBuilder sb = new StringBuilder();
        for (Map.Entry<String, String> entry : params.entrySet()) {
            if (entry.getValue() == null) {
                continue;
            }
            sb.append(entry.getKey());
            sb.append("=");
            sb.append(URLEncoder.encode(entry.getValue(), "UTF-8"));
            sb.append("&");
        }
        if (sb.length() > 0) {
            sb.deleteCharAt(sb.length() - 1);
        }
        conn.getOutputStream().write(sb.toString().getBytes("UTF-8"));
        conn.getOutputStream().flush();
        conn.getOutputStream().close();

        InputStream is = conn.getInputStream();
        Reader reader = new InputStreamReader(is, "UTF-8");
        StringBuilder sb2 = new StringBuilder();
        char[] buf = new char[1024];
        int len = 0;
        while ((len = reader.read(buf)) != -1) {
            sb2.append(buf, 0, len);
        }
        reader.close();
        is.close();
        conn.disconnect();

        return sb2.toString().getBytes("UTF-8");
    }


    public SqlQueryDto Query(SqlQueryForm form) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("project", String.valueOf(form.getProject()));
        payload.put("iid", form.getIid());
        payload.put("sql", form.getSql());
        payload.put("owner", form.getOwner());
        payload.put("schema", form.getSchema());
        payload.put("page_index", String.valueOf(form.getPageIndex()));
        payload.put("page_size", String.valueOf(form.getPageSize()));
        payload.put("need_total", String.valueOf(form.isNeedTotal()));
        byte[] response = this.sendRequest("POST", "/sql/query", payload);
        return JsonUtil.fromJson(new String(response, "UTF-8"), SqlQueryDto.class);
    }

    public UserDto GetUserInfo(int id) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("id", String.valueOf(id));
        byte[] response = this.sendRequest("POST", "/user/info", payload);
        return JsonUtil.fromJson(new String(response, "UTF-8"), UserDto.class);
    }

    public UserPageDto GetUserPage(int pageIndex, int pageSize) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("page_index", String.valueOf(pageIndex));
        payload.put("page_size", String.valueOf(pageSize));
        byte[] response = this.sendRequest("POST", "/user/page", payload);
        return JsonUtil.fromJson(new String(response, "UTF-8"), UserPageDto.class);
    }

    public String AddUser(UserAddForm user) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("name", user.getName());
        payload.put("user_name", user.getUserName());
        payload.put("pwd", user.getPwd());
        payload.put("expire", String.valueOf(user.getExpire()));
        payload.put("role", String.valueOf(user.getRole()));
        payload.put("email", user.getEmail());
        payload.put("telno", user.getTelno());
        payload.put("roles_str", user.getRolesStr());
        payload.put("group_name", user.getGroupName());
        byte[] response = this.sendRequest("POST", "/user/add", payload);
        return new String(response, "UTF-8");
    }

    public String UpdUser(UserUpdForm user) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("id", String.valueOf(user.getId()));
        payload.put("name", user.getName());
        payload.put("user_name", user.getUserName());
        payload.put("expire", String.valueOf(user.getExpire()));
        payload.put("role", String.valueOf(user.getRole()));
        payload.put("sex", String.valueOf(user.getSex()));
        payload.put("email", user.getEmail());
        payload.put("telno", user.getTelno());
        payload.put("roles_str", user.getRolesStr());
        payload.put("group_name", user.getGroupName());
        byte[] response = this.sendRequest("POST", "/user/upd", payload);
        return new String(response, "UTF-8");
    }

    public String RemoveUser(int id) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("id", String.valueOf(id));
        byte[] response = this.sendRequest("POST", "/user/remove", payload);
        return new String(response, "UTF-8");
    }

    public String StateUser(int id, int state) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("id", String.valueOf(id));
        payload.put("state", String.valueOf(state));
        byte[] response = this.sendRequest("POST", "/user/state", payload);
        return new String(response, "UTF-8");
    }

    public InstanceDto GetInstanceInfo(String id) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("id", id);
        byte[] response = this.sendRequest("POST", "/instance/info", payload);
        return JsonUtil.fromJson(new String(response, "UTF-8"), InstanceDto.class);
    }

    public InstancePageDto GetInstancePage(int pageIndex, int pageSize) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("page_index", String.valueOf(pageIndex));
        payload.put("page_size", String.valueOf(pageSize));
        byte[] response = this.sendRequest("POST", "/instance/page", payload);
        return JsonUtil.fromJson(new String(response, "UTF-8"), InstancePageDto.class);
    }

    public String AddInstance(InstanceAddForm instance) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("name", instance.getName());
        payload.put("user", instance.getUser());
        payload.put("pass", instance.getPass());
        payload.put("ip", instance.getIp());
        payload.put("port", instance.getPort());
        payload.put("db", instance.getDb());
        payload.put("con_type", instance.getCon_type());
        payload.put("service_type", instance.getService_type());
        payload.put("user_type", instance.getUser_type());
        payload.put("oper_type", String.valueOf(instance.getOper_type()));
        payload.put("db_type", instance.getDb_type());
        payload.put("asyn_export", String.valueOf(instance.getAsyn_export()));
        payload.put("asyn_query", String.valueOf(instance.getAsyn_query()));
        payload.put("asyn_sql", String.valueOf(instance.getAsyn_sql()));
        payload.put("page_query", String.valueOf(instance.getPage_query()));
        payload.put("config", instance.getConfig());
        payload.put("pool_min", String.valueOf(instance.getPool_min()));
        payload.put("pool_max", String.valueOf(instance.getPool_max()));
        payload.put("max_row", String.valueOf(instance.getMax_row()));
        payload.put("ssh_use", String.valueOf(instance.isSsh_use()));
        payload.put("ssh_server", instance.getSsh_server());
        payload.put("ssh_port", instance.getSsh_port());
        payload.put("ssh_user", instance.getSsh_user());
        payload.put("ssh_password", instance.getSsh_password());
        payload.put("ssh_auth", String.valueOf(instance.getSsh_auth()));
        payload.put("tx_timeout", String.valueOf(instance.getTx_timeout()));

        byte[] response = this.sendRequest("POST", "/instance/add", payload);
        return new String(response, "UTF-8");
    }

    public String UpdInstance(InstanceUpdForm instance) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("id", instance.getId());
        payload.put("name", instance.getName());
        payload.put("user", instance.getUser());
        payload.put("pass", instance.getPass());
        payload.put("ip", instance.getIp());
        payload.put("port", instance.getPort());
        payload.put("db", instance.getDb());
        payload.put("con_type", instance.getCon_type());
        payload.put("service_type", instance.getService_type());
        payload.put("user_type", instance.getUser_type());
        payload.put("oper_type", String.valueOf(instance.getOper_type()));
        payload.put("db_type", instance.getDb_type());
        payload.put("asyn_export", String.valueOf(instance.getAsyn_export()));
        payload.put("asyn_query", String.valueOf(instance.getAsyn_query()));
        payload.put("asyn_sql", String.valueOf(instance.getAsyn_sql()));
        payload.put("page_query", String.valueOf(instance.getPage_query()));
        payload.put("config", instance.getConfig());
        payload.put("pool_min", String.valueOf(instance.getPool_min()));
        payload.put("pool_max", String.valueOf(instance.getPool_max()));
        payload.put("max_row", String.valueOf(instance.getMax_row()));
        payload.put("ssh_use", String.valueOf(instance.isSsh_use()));
        payload.put("ssh_server", instance.getSsh_server());
        payload.put("ssh_port", instance.getSsh_port());
        payload.put("ssh_user", instance.getSsh_user());
        payload.put("ssh_password", instance.getSsh_password());
        payload.put("ssh_auth", String.valueOf(instance.getSsh_auth()));
        payload.put("tx_timeout", String.valueOf(instance.getTx_timeout()));

        byte[] response = this.sendRequest("POST", "/instance/upd", payload);
        return new String(response, "UTF-8");
    }

    public String RemoveInstance(String id) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("id", id);
        byte[] response = this.sendRequest("POST", "/instance/remove", payload);
        return new String(response, "UTF-8");
    }

    public String StateInstance(String id, int state) throws Exception {
        Map<String, String> payload = new HashMap<>();
        payload.put("id", id);
        payload.put("state", String.valueOf(state));
        byte[] response = this.sendRequest("POST", "/instance/state", payload);
        return new String(response, "UTF-8");
    }

    public static void main(String[] args) throws Exception {
        Sqldev sqldev = new Sqldev("http://127.0.0.1:8080/openapi/v1", "a9967fcfa669c40c5df7db31a33f70d4", "5414aabf-db54-4f45-b7e1-1d99aa10e47c");
        InstancePageDto a = sqldev.GetInstancePage(0, 10);
        System.out.println(a);

        sqldev.StateInstance("3afcc6c7a6e3b5cbb9393094746ba948",0);
    }
}
