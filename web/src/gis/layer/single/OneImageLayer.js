import ImageLayer from '../ImageLayer';

/**
 * 一张图片图层
 * @author tengge / https://github.com/tengge1
 * @param {*} globe 地球
 */
function OneImageLayer(globe) {
    ImageLayer.call(this, globe);
}

OneImageLayer.prototype = Object.create(ImageLayer.prototype);
OneImageLayer.prototype.constructor = OneImageLayer;

OneImageLayer.prototype.get = function (aabb) { // eslint-disable-line
    // TODO: 根据包围盒返回一张图片
};

export default OneImageLayer;