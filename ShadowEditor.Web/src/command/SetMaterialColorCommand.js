import Command from './Command';

/**
 * 设置材质颜色命令
 * @author dforrer / https://github.com/dforrer
 * Developed as part of a project at University of Applied Sciences and Arts Northwestern Switzerland (www.fhnw.ch)
 * @param object THREE.Object3D
 * @param attributeName string
 * @param newValue integer representing a hex color value
 * @constructor
 */
function SetMaterialColorCommand(object, attributeName, newValue) {
	Command.call(this);

	this.type = 'SetMaterialColorCommand';
	this.name = L_SET_MATERIAL + '.' + attributeName;
	this.updatable = true;

	this.object = object;
	this.attributeName = attributeName;
	this.oldValue = (object !== undefined) ? this.object.material[this.attributeName].getHex() : undefined;
	this.newValue = newValue;
};

SetMaterialColorCommand.prototype = Object.create(Command.prototype);

Object.assign(SetMaterialColorCommand.prototype, {
	constructor: SetMaterialColorCommand,

	execute: function () {
		this.object.material[this.attributeName].setHex(this.newValue);
		app.call('materialChanged', this, this.object.material);
	},

	undo: function () {
		this.object.material[this.attributeName].setHex(this.oldValue);
		app.call('materialChanged', this, this.object.material);
	},

	update: function (cmd) {
		this.newValue = cmd.newValue;
	},

	toJSON: function () {
		var output = Command.prototype.toJSON.call(this);

		output.objectUuid = this.object.uuid;
		output.attributeName = this.attributeName;
		output.oldValue = this.oldValue;
		output.newValue = this.newValue;

		return output;
	},

	fromJSON: function (json) {
		Command.prototype.fromJSON.call(this, json);

		this.object = this.editor.objectByUuid(json.objectUuid);
		this.attributeName = json.attributeName;
		this.oldValue = json.oldValue;
		this.newValue = json.newValue;
	}
});

export default SetMaterialColorCommand;
