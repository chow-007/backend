INSERT INTO `videos` VALUES (1, '1#', 'http://vjs.zencdn.net/v/oceans.mp4');
INSERT INTO `videos` VALUES (2, '2#', 'rtsp://211.139.194.251:554/live/2/13E6330A31193128/5iLd2iNl5nQ2s8r8.sdp');
INSERT INTO `videos` VALUES (3, '3#', 'http://vjs.zencdn.net/v/oceans.mp4');
INSERT INTO `videos` VALUES (4, '1#', 'http://vjs.zencdn.net/v/oceans.mp4');


INSERT INTO `permissions` VALUES (103, '权限管理', 0, 'per', 0, 3);
INSERT INTO `permissions` VALUES (125, '用户管理', 0, 'user', 0, 2);
INSERT INTO `permissions` VALUES (127, '数据中心', 0, 'per', 0, 1);
INSERT INTO `permissions` VALUES (140, '视频监控', 0, 'df', 0, 4);
INSERT INTO `permissions` VALUES (141, '用户列表', 125, 'users', 1, 1);
INSERT INTO `permissions` VALUES (142, '角色列表', 103, 'roles', 1, 1);
INSERT INTO `permissions` VALUES (143, '权限列表', 103, 'rights', 1, 2);
INSERT INTO `permissions` VALUES (144, '实时数据', 127, 'realtime', 1, 1);
INSERT INTO `permissions` VALUES (145, '历史数据', 127, 'history', 1, 2);
INSERT INTO `permissions` VALUES (146, '服务器', 127, 'server', 1, 3);
INSERT INTO `permissions` VALUES (147, '摄像头', 140, 'camera', 1, 1);
INSERT INTO `permissions` VALUES (148, '视频列表', 140, 'videos', 1, 2);


