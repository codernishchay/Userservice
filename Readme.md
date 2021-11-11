
 

## User Service 
### this app has 4 endpoints : 

### create : 
        using this method we can create a user to database            
### delete 
        we can  delete a user using the userid 
### update 
        using the user id we can update the user data 
### get 
       this endpoint will return all the users in the database 


## User Model 
    
    type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	DOB         string             `json:"dob" bson:"dob"`
	Address     string             `json:"address" bson:"address"`
	Description string             `json:"description" bson:"description"`
}




# Hard Machine

```jsx
[sudo] password for nishi: 
Starting Nmap 7.92 ( https://nmap.org ) at 2021-09-23 09:48 EDT
Stats: 0:00:03 elapsed; 0 hosts completed (1 up), 1 undergoing SYN Stealth Scan
SYN Stealth Scan Timing: About 19.45% done; ETC: 09:48 (0:00:12 remaining)
Nmap scan report for 10.10.11.103
Host is up (1.2s latency).
Not shown: 998 closed tcp ports (reset)
PORT   STATE SERVICE VERSION
22/tcp open  ssh     OpenSSH 8.2p1 Ubuntu 4ubuntu0.3 (Ubuntu Linux; protocol 2.0)
| ssh-hostkey: 
|   3072 36:aa:93:e4:a4:56:ab:39:86:66:bf:3e:09:fa:eb:e0 (RSA)
|   256 11:fb:e9:89:2e:4b:66:40:7b:6b:01:cf:f2:f2:ee:ef (ECDSA)
|_  256 77:56:93:6e:5f:ea:e2:ad:b0:2e:cf:23:9d:66:ed:12 (ED25519)
80/tcp open  http    Apache httpd 2.4.41
|_http-title: Did not follow redirect to http://developer.htb/
|_http-server-header: Apache/2.4.41 (Ubuntu)
Service Info: Host: developer.htb; OS: Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 30.11 seconds
```

using directory bustering 

```jsx
/_admin               (Status: 301) [Size: 0] [--> /_admin/]
/~admin               (Status: 301) [Size: 0] [--> /~admin/]
/~sysadmin            (Status: 301) [Size: 0] [--> /~sysadmin/]
/admin                (Status: 301) [Size: 0] [--> /admin/]    
/admin-admin          (Status: 301) [Size: 0] [--> /admin-admin/]
/aspadmin             (Status: 301) [Size: 0] [--> /aspadmin/]   
/axis2-admin          (Status: 301) [Size: 0] [--> /axis2-admin/]
/axis-admin           (Status: 301) [Size: 0] [--> /axis-admin/] 
/banneradmin          (Status: 301) [Size: 0] [--> /banneradmin/]
/bbadmin              (Status: 301) [Size: 0] [--> /bbadmin/]    
/bigadmin             (Status: 301) [Size: 0] [--> /bigadmin/]   
/ccp14admin           (Status: 301) [Size: 0] [--> /ccp14admin/] 
/cmsadmin             (Status: 301) [Size: 0] [--> /cmsadmin/]   
/contact              (Status: 301) [Size: 0] [--> /contact/]    
/cpadmin              (Status: 301) [Size: 0] [--> /cpadmin/]    
/dashboard            (Status: 301) [Size: 0] [--> /dashboard/]  
/dbadmin              (Status: 301) [Size: 0] [--> /dbadmin/]    
/dh_phpmyadmin        (Status: 301) [Size: 0] [--> /dh_phpmyadmin/]
/directadmin          (Status: 301) [Size: 0] [--> /directadmin/]  
/e107_admin           (Status: 301) [Size: 0] [--> /e107_admin/]   
/ezsqliteadmin        (Status: 301) [Size: 0] [--> /ezsqliteadmin/]
/fileadmin            (Status: 301) [Size: 0] [--> /fileadmin/]    
/globes_admin         (Status: 301) [Size: 0] [--> /globes_admin/] 
/hpwebjetadmin        (Status: 301) [Size: 0] [--> /hpwebjetadmin/]
/iisadmin             (Status: 301) [Size: 0] [--> /iisadmin/]     
/index_admin          (Status: 301) [Size: 0] [--> /index_admin/]  
/Indy_admin           (Status: 301) [Size: 0] [--> /Indy_admin/]   
/indy_admin           (Status: 301) [Size: 0] [--> /indy_admin/]   
/INSTALL_admin        (Status: 301) [Size: 0] [--> /INSTALL_admin/]
/irc-macadmin         (Status: 301) [Size: 0] [--> /irc-macadmin/] 
/listadmin            (Status: 301) [Size: 0] [--> /listadmin/]    
/loginadmin           (Status: 301) [Size: 0] [--> /loginadmin/]   
/logo_sysadmin        (Status: 301) [Size: 0] [--> /logo_sysadmin/]
/macadmin             (Status: 301) [Size: 0] [--> /macadmin/]     
/media                (Status: 301) [Size: 314] [--> http://developer.htb/media/]
/myadmin              (Status: 301) [Size: 0] [--> /myadmin/]                    
/navsiteadmin         (Status: 301) [Size: 0] [--> /navsiteadmin/]               
/newadmin             (Status: 301) [Size: 0] [--> /newadmin/]                   
/newsadmin            (Status: 301) [Size: 0] [--> /newsadmin/]                  
/openvpnadmin         (Status: 301) [Size: 0] [--> /openvpnadmin/]               
/pgadmin              (Status: 301) [Size: 0] [--> /pgadmin/]                    
/phpadmin             (Status: 301) [Size: 0] [--> /phpadmin/]                   
/phpldapadmin         (Status: 301) [Size: 0] [--> /phpldapadmin/]               
/phpmyadmin           (Status: 301) [Size: 0] [--> /phpmyadmin/]                 
/phppgadmin           (Status: 301) [Size: 0] [--> /phppgadmin/]                 
/profile              (Status: 301) [Size: 0] [--> /profile/]                    
/pureadmin            (Status: 301) [Size: 0] [--> /pureadmin/]                  
/resin-admin          (Status: 301) [Size: 0] [--> /resin-admin/]                
/server-status        (Status: 403) [Size: 278]                                  
/shopadmin            (Status: 301) [Size: 0] [--> /shopadmin/]                  
/siteadmin            (Status: 301) [Size: 0] [--> /siteadmin/]                  
/sohoadmin            (Status: 301) [Size: 0] [--> /sohoadmin/]                  
/sqladmin             (Status: 301) [Size: 0] [--> /sqladmin/]                   
/sql-admin            (Status: 301) [Size: 0] [--> /sql-admin/]                  
/sshadmin             (Status: 301) [Size: 0] [--> /sshadmin/]                   
/staradmin            (Status: 301) [Size: 0] [--> /staradmin/]                  
/static               (Status: 301) [Size: 315] [--> http://developer.htb/static/]
/sysadmin             (Status: 301) [Size: 0] [--> /sysadmin/]                    
/sys-admin            (Status: 301) [Size: 0] [--> /sys-admin/]                   
/system_admin         (Status: 301) [Size: 0] [--> /system_admin/]                
/system-admin         (Status: 301) [Size: 0] [--> /system-admin/]                
/topicadmin           (Status: 301) [Size: 0] [--> /topicadmin/]                  
/ur-admin             (Status: 301) [Size: 0] [--> /ur-admin/]                    
/useradmin            (Status: 301) [Size: 0] [--> /useradmin/]                   
/vmailadmin           (Status: 301) [Size: 0] [--> /vmailadmin/]                  
/vsadmin              (Status: 301) [Size: 0] [--> /vsadmin/]                     
/wbsadmin             (Status: 301) [Size: 0] [--> /wbsadmin/]                    
/webadmin             (Status: 301) [Size: 0] [--> /webadmin/]                    
/wizmysqladmin        (Status: 301) [Size: 0] [--> /wizmysqladmin/]               
/wp-admin             (Status: 301) [Size: 0] [--> /wp-admin/]                    
Progress: 4578 / 4626 (98.96%)
```

directories i got by gobuster 

all of them are 301 ?