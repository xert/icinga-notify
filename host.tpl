{{define "Subject"}}
{{.NOTIFICATIONTYPE}} - {{.HOSTDISPLAYNAME}} is {{.HOSTSTATE}}
{{end}}

{{define "Content"}}
***** Icinga  *****

Notification Type: {{.NOTIFICATIONTYPE}}

Host: {{.HOSTALIAS}}
Address: {{.HOSTADDRESS}}
State: {{.HOSTSTATE}}

Date/Time: {{.LONGDATETIME}}

Additional Info: {{.HOSTOUTPUT}}

Comment: [{{.NOTIFICATIONAUTHORNAME}}] {{.NOTIFICATIONCOMMENT}}
{{end}}
{{define "Plain"}}
***** Icinga  *****

Notification Type: {{.NOTIFICATIONTYPE}}

Host: {{.HOSTALIAS}}
Address: {{.HOSTADDRESS}}
State: {{.HOSTSTATE}}

Date/Time: {{.LONGDATETIME}}

Additional Info: {{.HOSTOUTPUT}}

Comment: [{{.NOTIFICATIONAUTHORNAME}}] {{.NOTIFICATIONCOMMENT}}
{{end}}
