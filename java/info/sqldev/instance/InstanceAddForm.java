package info.sqldev.instance;

public class InstanceAddForm {
    private String name;
    private String user;
    private String pass;
    private String ip;
    private String port;
    private String db;
    private String con_type;
    private String service_type;
    private String user_type;
    private int oper_type;
    private String db_type;
    private int asyn_export;
    private int asyn_query;
    private int asyn_sql;
    private int page_query;
    private String config;
    private int pool_min;
    private int pool_max;
    private int max_row;
    private boolean ssh_use;
    private String ssh_server;
    private String ssh_port;
    private String ssh_user;
    private String ssh_password;
    private int ssh_auth;
    private long tx_timeout;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getUser() {
        return user;
    }

    public void setUser(String user) {
        this.user = user;
    }

    public String getPass() {
        return pass;
    }

    public void setPass(String pass) {
        this.pass = pass;
    }

    public String getIp() {
        return ip;
    }

    public void setIp(String ip) {
        this.ip = ip;
    }

    public String getPort() {
        return port;
    }

    public void setPort(String port) {
        this.port = port;
    }

    public String getDb() {
        return db;
    }

    public void setDb(String db) {
        this.db = db;
    }

    public String getCon_type() {
        return con_type;
    }

    public void setCon_type(String con_type) {
        this.con_type = con_type;
    }

    public String getService_type() {
        return service_type;
    }

    public void setService_type(String service_type) {
        this.service_type = service_type;
    }

    public String getUser_type() {
        return user_type;
    }

    public void setUser_type(String user_type) {
        this.user_type = user_type;
    }

    public int getOper_type() {
        return oper_type;
    }

    public void setOper_type(int oper_type) {
        this.oper_type = oper_type;
    }

    public String getDb_type() {
        return db_type;
    }

    public void setDb_type(String db_type) {
        this.db_type = db_type;
    }

    public int getAsyn_export() {
        return asyn_export;
    }

    public void setAsyn_export(int asyn_export) {
        this.asyn_export = asyn_export;
    }

    public int getAsyn_query() {
        return asyn_query;
    }

    public void setAsyn_query(int asyn_query) {
        this.asyn_query = asyn_query;
    }

    public int getAsyn_sql() {
        return asyn_sql;
    }

    public void setAsyn_sql(int asyn_sql) {
        this.asyn_sql = asyn_sql;
    }

    public int getPage_query() {
        return page_query;
    }

    public void setPage_query(int page_query) {
        this.page_query = page_query;
    }

    public String getConfig() {
        return config;
    }

    public void setConfig(String config) {
        this.config = config;
    }

    public int getPool_min() {
        return pool_min;
    }

    public void setPool_min(int pool_min) {
        this.pool_min = pool_min;
    }

    public int getPool_max() {
        return pool_max;
    }

    public void setPool_max(int pool_max) {
        this.pool_max = pool_max;
    }

    public int getMax_row() {
        return max_row;
    }

    public void setMax_row(int max_row) {
        this.max_row = max_row;
    }

    public boolean isSsh_use() {
        return ssh_use;
    }

    public void setSsh_use(boolean ssh_use) {
        this.ssh_use = ssh_use;
    }

    public String getSsh_server() {
        return ssh_server;
    }

    public void setSsh_server(String ssh_server) {
        this.ssh_server = ssh_server;
    }

    public String getSsh_port() {
        return ssh_port;
    }

    public void setSsh_port(String ssh_port) {
        this.ssh_port = ssh_port;
    }

    public String getSsh_user() {
        return ssh_user;
    }

    public void setSsh_user(String ssh_user) {
        this.ssh_user = ssh_user;
    }

    public String getSsh_password() {
        return ssh_password;
    }

    public void setSsh_password(String ssh_password) {
        this.ssh_password = ssh_password;
    }

    public int getSsh_auth() {
        return ssh_auth;
    }

    public void setSsh_auth(int ssh_auth) {
        this.ssh_auth = ssh_auth;
    }

    public long getTx_timeout() {
        return tx_timeout;
    }

    public void setTx_timeout(long tx_timeout) {
        this.tx_timeout = tx_timeout;
    }
}
