import Vue from 'vue';
import VueAMap from 'vue-amap';
Vue.use(VueAMap);

if (!Vue.prototype.$isServer) {
    VueAMap.initAMapApiLoader({
        key: '2a00368056ce869e22681f75eb121345',
        plugin: ['AMap.Geolocation', 'AMap.Marker', 'AMap.ToolBar', 'AMap.Circle', 'AMap.PolyLine'],
        uiVersion: '1.0',
        v: '1.4.15'
    });
}