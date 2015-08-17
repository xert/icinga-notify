{{define "Subject"}}{{env "NOTIFICATIONTYPE"}} - {{env "HOSTDISPLAYNAME"}} - {{env "SERVICEDISPLAYNAME"}} is {{env "HOSTSTATE"}}{{end}}

{{define "Content"}}***** OKO  *****

Notification Type: {{env "NOTIFICATIONTYPE"}}

Service: {{env "SERVICEDESC"}}
Host: {{env "HOSTALIAS"}}
Address: {{env "HOSTADDRESS"}}
State: {{env "SERVICESTATE"}}

Date/Time: {{env "LONGDATETIME"}}

Additional Info: {{env "SERVICEOUTPUT"}}

Comment: [{{env "NOTIFICATIONAUTHORNAME"}}] {{env "NOTIFICATIONCOMMENT"}}
{{end}}
