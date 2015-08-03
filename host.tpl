{{define "Subject"}}
{{env "NOTIFICATIONTYPE"}} - {{env "HOSTDISPLAYNAME"}} is {{env "HOSTSTATE"}}
{{end}}

{{define "Content"}}
***** OKO-1  *****

Notification Type: {{env "NOTIFICATIONTYPE"}}

Host: {{env "HOSTALIAS"}}
Address: {{env "HOSTADDRESS"}}
State: {{env "HOSTSTATE"}}

Date/Time: {{env "LONGDATETIME"}}

Additional Info: {{env "HOSTOUTPUT"}}

Comment: [{{env "NOTIFICATIONAUTHORNAME"}}] {{env "NOTIFICATIONCOMMENT"}}
{{end}}
